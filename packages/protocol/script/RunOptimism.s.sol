// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/console.sol";
import {ScriptPlus} from "./ScriptPlus.s.sol";
import {AaveV3Optimism} from "../src/providers/optimism/AaveV3Optimism.sol";
import {DForceOptimism} from "../src/providers/optimism/DForceOptimism.sol";
import {WePiggyOptimism} from "../src/providers/optimism/WePiggyOptimism.sol";

contract RunOptimism is ScriptPlus {
  AaveV3Optimism aaveV3;
  DForceOptimism dforce;
  WePiggyOptimism wePiggy;

  function setUp() public {
    setUpOn();
  }

  function run() public {
    vm.startBroadcast(deployer);

    setOrDeployChief(false);
    setOrDeployConnextRouter(false);
    setOrDeployEverestOracle(false);
    setOrDeployBorrowingVaultFactory(false, false);
    setOrDeployYieldVaultFactory(false, false);
    setOrDeployAddrMapper(false);
    setOrDeployFlasherBalancer(false);
    setOrDeployRebalancer(false);

    _setLendingProviders();

    if (chief.allowedVaultFactory(address(factory))) {
      deployBorrowingVaults();
      /*setBorrowingVaults();*/
    }

    if (chief.allowedVaultFactory(address(yieldFactory))) {
      deployYieldVaults();
    }

    /*setVaultNewRating("BorrowingVault-WETHUSDC", 75);*/
    /*rebalanceVault("BorrowingVault-WETHUSDC", compound, aaveV3);*/
    /*rebalanceYieldVaults();*/
    /*rebalanceBorrowingVaults();*/

    // If setting all routers at once, call after deploying all chians
    /*setConnextReceivers();*/

    /*upgradeBorrowingImpl(false);*/

    /*bytes memory constructorArgs = abi.encode(getAddress("ConnextRouter"));*/
    /*verifyContract("ConnextHandler", constructorArgs);*/

    vm.stopBroadcast();
  }

  function _setLendingProviders() internal {
    aaveV3 = AaveV3Optimism(getAddress("Aave_V3_Optimism_Emode"));
    /*aaveV3 = new AaveV3Optimism();*/
    /*saveAddress("Aave_V3_Optimism_Emode", address(aaveV3));*/
    setOrdeployAaveEModeHelper(false);

    dforce = DForceOptimism(getAddress("DForce_Optimism"));
    /*dforce = new DForceOptimism();*/
    /*saveAddress("DForce_Optimism", address(dforce));*/

    wePiggy = WePiggyOptimism(getAddress("We_Piggy_Optimism"));
    /*wePiggy = new WePiggyOptimism();*/
    /*saveAddress("We_Piggy_Optimism", address(wePiggy));*/
  }
}
