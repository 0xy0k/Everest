module Cube where

import Prelude

import Control.Monad.Aff (Aff)
import Control.Monad.Aff.AVar (AVAR)
import Control.Monad.Eff.Console (log, CONSOLE)
import DOM (DOM)
import DOM.HTML (window) as DOM
import DOM.HTML.Types (htmlDocumentToDocument) as DOM
import DOM.HTML.Window (document) as DOM
import Data.Array (dropEnd, length, mapWithIndex, snoc, (!!))
import Data.Maybe (Maybe(..), fromMaybe)
import Data.Tuple (Tuple(..))
import Global as G
import Halogen as H
import Halogen.HTML as HH
import Halogen.HTML.Events as HE
import Halogen.HTML.Properties as HP
import Halogen.Query.EventSource as ES
import Keyboard as K
import Math (cos, sin)
import Svg.Attributes as SA
import Svg.Elements as SE

-- Core Types
type Distance = Number

type Angle = Number

type Point2D =
  { x :: Distance
  , y :: Distance
  }

type Point3D =
  { x :: Distance
  , y :: Distance
  , z :: Distance
  }

type Edge = Tuple Int Int
type Face =
  { p1 :: Int
  , p2 :: Int
  , p3 :: Int
  , p4 :: Int
  , r :: Int
  , g :: Int
  , b :: Int
  }

type Shape =
  { vertices :: Array Point3D
  , edges :: Array Edge
  , faces :: Array Face
  }

type Angle3D =
  { xa :: Angle
  , ya :: Angle
  , za :: Angle
  }

type AngVelocity3D = Angle3D -- velocity = angle/sec

type RotatingShape =
  { shape :: Shape
  , angVel :: AngVelocity3D
  , forward :: Boolean
  , valspeed :: Number
  , id :: Int
  }

data Axis = X | Y | Z

-- Model / State
type State = {
  allshapes :: Array RotatingShape,
  cuberate :: Number
}

-- Values

viewBoxSize :: Number
viewBoxSize = 600.0

viewCenter :: Point2D
viewCenter =
  { x: viewBoxSize / 2.0
  , y: viewBoxSize / 2.0
  }

frameRate :: Number
frameRate = 200.0

oneDegInRad :: Angle
oneDegInRad = 0.01745329255

tenDegInRad :: Angle
tenDegInRad = oneDegInRad * 10.0

accelerateBy :: Number
accelerateBy = oneDegInRad * 50.0

dampenPercent :: Number
dampenPercent = 1.0 - (0.9 / frameRate) -- 10% per second

initShape :: RotatingShape
initShape = { shape:
        { vertices:
            [ { x:  100.0, y:  100.0, z:  100.0 }
            , { x: -100.0, y:  100.0, z:  100.0 }
            , { x:  100.0, y: -100.0, z:  100.0 }
            , { x: -100.0, y: -100.0, z:  100.0 }
            , { x:  100.0, y:  100.0, z: -100.0 }
            , { x: -100.0, y:  100.0, z: -100.0 }
            , { x:  100.0, y: -100.0, z: -100.0 }
            , { x: -100.0, y: -100.0, z: -100.0 }
            ]
        , edges:
            [ Tuple 0 1
            , Tuple 0 2
            , Tuple 0 4
            , Tuple 1 5
            , Tuple 1 3
            , Tuple 2 3
            , Tuple 2 6
            , Tuple 4 5
            , Tuple 4 6
            , Tuple 3 7
            , Tuple 6 7
            , Tuple 5 7
            ]
        , faces:
            [ {p1: 1, p2: 5, p3: 7, p4: 3, r: 255, g: 0, b: 0}
            , {p1: 5, p2: 4, p3: 6, p4: 7, r: 0, g: 255, b: 0}
            , {p1: 7, p2: 6, p3: 2, p4: 3, r: 0, g: 0, b: 255}
            , {p1: 1, p2: 3, p3: 2, p4: 0, r: 255, g: 255, b: 0}
            , {p1: 1, p2: 0, p3: 4, p4: 5, r: 255, g: 0, b: 255}
            , {p1: 2, p2: 6, p3: 4, p4: 0, r: 0, g: 255, b: 255}
            ]
        }
    , angVel:
        { xa: tenDegInRad
        , ya: tenDegInRad
        , za: tenDegInRad
        }
    , forward: true
    , valspeed: 1.0
    , id: 0
    }

initCube :: State
initCube = 
  { 
    allshapes: [initShape],
    cuberate: 1.0
  }

-- Events
data Query a
  = Tick a
  | IncAngVelocity Axis a
  | AddCube a
  | RemoveCube a
  | ReverseDirection Int a
  | IncVelocity Int a
  | DecVelocity Int a
  | AdjustCubeSize String a
  | Init a
  | HandleKey K.KeyboardEvent (H.SubscribeStatus -> a)

-------------------- UPDATE / REDUCERS --------------------

cubes :: forall eff. H.Component HH.HTML Query Unit Unit (Aff (console :: CONSOLE, dom :: DOM, avar :: AVAR, keyboard :: K.KEYBOARD | eff))
cubes =
  H.component
    { initialState: const initialState
    , render
    , eval
    , receiver: const Nothing
    }
  where
    initialState :: State
    initialState = initCube

    render :: State -> H.ComponentHTML Query
    render = renderView

    eval :: Query ~> H.ComponentDSL State Query Unit (Aff (console :: CONSOLE, dom :: DOM, avar :: AVAR, keyboard :: K.KEYBOARD | eff))
    eval = case _ of
      Tick next -> do
        cube <- H.get
        let newCube = updateCube cube
        H.put newCube
        H.liftEff $ log "tick"
        pure next

      IncAngVelocity axis next -> do
        cube <- H.get
        let newCube = updateIncAngVelocity cube axis
        H.modify
          (\c -> newCube)
        pure next

      AddCube next -> do
        cube <- H.get
        let newCube = addNextCube cube
        H.modify
          (\c -> newCube)
        pure next

      RemoveCube next -> do
        cube <- H.get
        let newCube = removeLastCube cube
        H.modify
          (\c -> newCube)
        pure next

      ReverseDirection curshape next -> do
        cube <- H.get
        let newCube = updateReverseDirection cube curshape
        H.modify
          (\c -> newCube)
        pure next

      IncVelocity curshape next -> do
        cube <- H.get
        let newCube = updateVelocity cube curshape 1.5
        H.modify
          (\c -> newCube)
        pure next

      DecVelocity curshape next -> do
        cube <- H.get
        let newCube = updateVelocity cube curshape 0.5
        H.modify
          (\c -> newCube)
        pure next

      AdjustCubeSize rate next -> do
        cube <- H.get
        let newCube = updateCubeSize cube rate
        H.modify
          (\c -> newCube)
        pure next
      
      Init next -> do
        H.liftEff $ log "shortcut init"
        document <- H.liftEff $ DOM.window >>= DOM.document <#> DOM.htmlDocumentToDocument
        H.subscribe $ ES.eventSource' (K.onKeyUp document) (Just <<< H.request <<< HandleKey)
        pure next
      
      HandleKey e reply -> do
        case K.readKeyboardEvent e of
          info
            | info.keyCode == 88 -> do
                cube <- H.get
                H.modify (\c -> updateIncAngVelocity cube X)
                pure (reply H.Listening)
            | info.keyCode == 89 -> do
                cube <- H.get
                H.modify (\c -> updateIncAngVelocity cube Y)
                pure (reply H.Listening)
            | info.keyCode == 90 -> do
                cube <- H.get
                H.modify (\c -> updateIncAngVelocity cube Z)
                pure (reply H.Listening)
            | info.keyCode == 65 -> do
                cube <- H.get
                H.modify (\c -> addNextCube cube)
                pure (reply H.Listening)
            | info.keyCode == 82 -> do
                cube <- H.get
                H.modify (\c -> removeLastCube cube)
                pure (reply H.Listening)
            | otherwise -> do
                pure (reply H.Listening)

fromString ∷ String → Number
fromString = G.readFloat >>> check
  where
    check num = num

updateCubeSize :: State -> String -> State
updateCubeSize cube newsize = newCube
  where
    numcubesize = (fromString newsize) / 50.0
    allshapes = map (updateShapeSize cube.cuberate numcubesize) cube.allshapes
    newCube = {allshapes: allshapes, cuberate: numcubesize}

updateShapeSize :: Number -> Number -> RotatingShape -> RotatingShape
updateShapeSize orgrate newrate orgshape = newshape
  where
    {shape, angVel, forward, valspeed, id} = orgshape
    {vertices, edges, faces} = shape
    newvertices = map (updatevertices orgrate newrate) vertices
    newshape = {shape: {vertices: newvertices, edges: edges, faces: faces}, angVel: angVel, forward: forward, valspeed: valspeed, id: id}

updatevertices :: Number -> Number -> Point3D -> Point3D
updatevertices orgrate newrate {x, y, z} = {x: newx, y: newy, z:newz}
  where
    newx = x / orgrate * newrate
    newy = y / orgrate * newrate
    newz = z / orgrate * newrate

removeLastCube :: State -> State
removeLastCube cube = newCube
  where
    allshapes = dropEnd 1 cube.allshapes
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}

addNextCube :: State -> State
addNextCube cube = newCube
  where
    {vertices, edges, faces} = initShape.shape
    newvertices = map (updatevertices 1.0 cube.cuberate) vertices
    newshape = { shape: {vertices: newvertices, edges: edges, faces: faces}
    , angVel: initShape.angVel
    , forward: true
    , valspeed: 1.0
    , id: length cube.allshapes
    }
    allshapes = snoc cube.allshapes newshape
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}

updateVelocity :: State -> Int ->Number -> State
updateVelocity cube curshape acc = newCube
 where
    allshapes = map (updateShapeVelocity curshape acc) cube.allshapes
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}

updateShapeVelocity :: Int -> Number ->RotatingShape -> RotatingShape
updateShapeVelocity curshape acc shape = if curshape == shape.id then newshape else shape
  where
    newvalspeed = shape.valspeed * acc
    newshape = {shape: shape.shape, angVel: shape.angVel, forward: shape.forward, valspeed: newvalspeed, id: shape.id}

updateReverseDirection :: State -> Int -> State
updateReverseDirection cube curshape = newCube
 where
    allshapes = map (updateShapeReverseDirection curshape) cube.allshapes
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}

updateShapeReverseDirection :: Int ->RotatingShape -> RotatingShape
updateShapeReverseDirection curshape shape = if curshape == shape.id then newshape else shape
  where
    newforward = if shape.forward == true then false else true
    newshape = {shape: shape.shape, angVel: shape.angVel, forward: newforward, valspeed: shape.valspeed, id: shape.id}

updateIncAngVelocity :: State -> Axis -> State
updateIncAngVelocity cube axis = newCube
  where
    allshapes = map (updateShapeAngVelocity axis) cube.allshapes
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}

updateShapeAngVelocity :: Axis -> RotatingShape -> RotatingShape
updateShapeAngVelocity axis shape = case axis of
    X -> newShapeX
    Y -> newShapeY
    Z -> newShapeZ
    where
      {xa, ya, za} = shape.angVel
      newangVelX = {xa: xa + getAccelerateBy accelerateBy shape.forward shape.valspeed, ya: ya, za: za}
      newangVelY = {xa: xa, ya: ya + getAccelerateBy accelerateBy shape.forward shape.valspeed, za: za}
      newangVelZ = {xa: xa, ya: ya, za: za + getAccelerateBy accelerateBy shape.forward shape.valspeed}
      newShapeX = {shape: shape.shape, angVel: newangVelX, forward: shape.forward, valspeed: shape.valspeed, id: shape.id}
      newShapeY = {shape: shape.shape, angVel: newangVelY, forward: shape.forward, valspeed: shape.valspeed, id: shape.id}
      newShapeZ = {shape: shape.shape, angVel: newangVelZ, forward: shape.forward, valspeed: shape.valspeed, id: shape.id}

updateCube :: State -> State
updateCube cube = newCube
  where
    allshapes = map updateShapes cube.allshapes
    newCube = {allshapes: allshapes, cuberate: cube.cuberate}
      
updateShapes :: RotatingShape -> RotatingShape
updateShapes shape = updatedShape
 where
  angVel = shape.angVel
  {vertices, edges, faces} = shape.shape
  newShape =
    { edges: edges
    , faces: faces
    , vertices: rotateShape vertices (anglePerFrame angVel)
    }
  updatedShape = shape
    { angVel = dampenAngVelocity angVel
    , shape = newShape
    }

getAccelerateBy :: Number -> Boolean -> Number -> Number
getAccelerateBy acc flag valspeed =
  if flag == true then acc * valspeed else -acc * valspeed

rotateShape :: Array Point3D -> AngVelocity3D -> Array Point3D
rotateShape vertices ang =
  map (rotate ang) vertices

rotate :: AngVelocity3D -> Point3D -> Point3D
rotate { xa, ya, za } = rotateX xa >>> rotateY ya >>> rotateZ za
  where
    rotateX ang {x,y,z} = let Tuple ny nz = rotateInPlane y z ang in { x, y:ny, z:nz }
    rotateY ang {x,y,z} = let Tuple nx nz = rotateInPlane x z ang in { x:nx, y, z:nz }
    rotateZ ang {x,y,z} = let Tuple nx ny = rotateInPlane x y ang in { x:nx, y:ny, z }

    rotateInPlane :: Number -> Number -> Number -> Tuple Number Number
    rotateInPlane axis1 axis2 ang =
      Tuple (axis1 * cos(ang) - axis2 * sin(ang)) (axis2 * cos(ang) + axis1 * sin(ang))

anglePerFrame :: AngVelocity3D -> Angle3D
anglePerFrame {xa, ya, za} =
  { xa: xa / frameRate
  , ya: ya / frameRate
  , za: za / frameRate
  }

dampenAngVelocity :: AngVelocity3D -> AngVelocity3D
dampenAngVelocity {xa, ya, za} =
    { xa: dampen xa
    , ya: dampen ya
    , za: dampen za
    }
  where
    dampen :: Number -> Number
    dampen ang = ang * dampenPercent -- Basics.max 0 (ang-drpf)

-------------------- VIEW --------------------
renderView :: State -> H.ComponentHTML Query
renderView state = let
    shapes = state.allshapes
  in
    HH.div_
      [
        HH.div_
        [
          HH.text "Adjust Cube Size: ",
          HH.input
          [
            HP.type_ HP.InputRange
          , HP.min 0.0
          , HP.max 100.0
          , HP.value "50"
          , HE.onValueChange $ HE.input AdjustCubeSize
          ]
        ],

        HH.div_
        [ 
          renderButton "rotX++" (IncAngVelocity X)
          , renderButton "rotY++" (IncAngVelocity Y)
          , renderButton "rotZ++" (IncAngVelocity Z)
          , renderButton "Add Cube" (AddCube)
          , renderButton "Remove Cube" (RemoveCube)
        ],

        HH.h1_
          [ HH.text "Click X/Y/Z to rotate along X-axis/Y-axis/Z-axis" ],

        HH.ol_ $ map (\shape -> 
          HH.div_
            [
              HH.div_
                [
                  renderButton "Reverse" (ReverseDirection shape.id)
                  , renderButton "vel++" (IncVelocity shape.id)
                  , renderButton "vel--" (DecVelocity shape.id)
                ],
              SE.svg
                [ SA.viewBox 0.0 0.0 viewBoxSize viewBoxSize ]
                [ SE.g []
                  (drawCubes shape)
                ]
            ]
          ) shapes
      ]
 
  where
    renderButton label query =
      HH.button
        [ HP.title label
        , HE.onClick (HE.input_ query)
        ]
        [ HH.text label ]

    -- parallel projection
    project :: Point3D -> Point2D
    project p =
      { x: p.x + viewCenter.x
      , y: p.y + viewCenter.y
      }

    drawCubes :: RotatingShape -> Array (H.ComponentHTML Query)
    drawCubes shape = drawCube edges faces vert2Ds
      where
        {vertices, edges, faces} = shape.shape
        vert2Ds = map project vertices

    drawCube :: Array Edge -> Array Face -> Array Point2D -> Array (H.ComponentHTML Query)
    drawCube edges faces vert2Ds =
      drawEdges edges vert2Ds <> drawFaces faces vert2Ds <> drawVertices vert2Ds
      -- drawEdges edges vert2Ds

    drawFaces :: Array Face -> Array Point2D -> Array (H.ComponentHTML Query)
    drawFaces faces verts =
        map (\face -> drawFace (getPoint (verts !! face.p1)) (getPoint (verts !! face.p2)) (getPoint (verts !! face.p3)) (getPoint (verts !! face.p4)) face.r face.g face.b) faces

    drawFace :: Point2D -> Point2D -> Point2D -> Point2D -> Int -> Int -> Int -> H.ComponentHTML Query
    drawFace p1 p2 p3 p4 r g b = if visible then
        SE.path
          [ SA.d
            [ SA.Abs (SA.M p1.x p1.y)
            , SA.Abs (SA.L p2.x p2.y)
            , SA.Abs (SA.L p3.x p3.y)
            , SA.Abs (SA.L p4.x p4.y)
            , SA.Abs (SA.L p1.x p1.y)
            ]
          , SA.fill $ Just (SA.RGB r g b)
          , SA.stroke $ Just (SA.RGB 50 50 50)
          ]
      else
        SE.path []
      where
        square = p1.x * p2.y - p2.x * p1.y + p2.x * p3.y - p3.x * p2.y + p3.x * p4.y - p4.x * p3.y + p4.x * p1.y - p1.x * p4.y
        visible = if square > 0.0 then true else false

    drawEdges :: Array Edge -> Array Point2D -> Array (H.ComponentHTML Query)
    drawEdges edges verts = let
        connectedVerts = map (\(Tuple v1 v2) -> Tuple (verts !! v1) (verts !! v2)) edges
      in
        map (\(Tuple v1 v2) -> drawLine (getPoint v1) (getPoint v2)) connectedVerts

    getPoint :: Maybe Point2D -> Point2D
    getPoint maybePoint = let
       default = { x: 100.0, y: 100.0 }
      in
        fromMaybe default maybePoint

    drawLine :: Point2D -> Point2D -> H.ComponentHTML Query
    drawLine a b =
      SE.path
        [ SA.d
          [ SA.Abs (SA.M a.x a.y)
          , SA.Abs (SA.L b.x b.y)
          ]
        , SA.stroke $ Just (SA.RGB 50 50 50)
        ]

    drawVertices :: Array Point2D -> Array (H.ComponentHTML Query)
    drawVertices vert2Ds =
      mapWithIndex drawVertex vert2Ds

    drawVertex :: Int -> Point2D -> H.ComponentHTML Query
    drawVertex idx {x, y} = SE.g []
      [ SE.text
          [ SA.x $ x + 5.0
          , SA.y $ y - 5.0
          , SA.fill $ Just (SA.RGB 150 150 150)
          ]
          [ HH.text $ show idx ]
      , SE.circle
          [ SA.r 3.0
          , SA.cx x
          , SA.cy y
          , SA.fill $ Just (SA.RGB 100 100 100)
          ]
      ]

