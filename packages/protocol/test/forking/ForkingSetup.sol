// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/console.sol";
import {Test} from "forge-std/Test.sol";
import {TimelockController} from
  "openzeppelin-contracts/contracts/governance/TimelockController.sol";
import {LibSigUtils} from "../../src/libraries/LibSigUtils.sol";
import {BorrowingVault} from "../../src/vaults/borrowing/BorrowingVault.sol";
import {YieldVault} from "../../src/vaults/yields/YieldVault.sol";
import {EverestOracle} from "../../src/EverestOracle.sol";
import {MockOracle} from "../../src/mocks/MockOracle.sol";
import {MockERC20} from "../../src/mocks/MockERC20.sol";
import {SafeERC20} from "openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {Chief} from "../../src/Chief.sol";
import {IVault} from "../../src/interfaces/IVault.sol";
import {IRouter} from "../../src/interfaces/IRouter.sol";
import {IVaultPermissions} from "../../src/interfaces/IVaultPermissions.sol";
import {ILendingProvider} from "../../src/interfaces/ILendingProvider.sol";
import {CoreRoles} from "../../src/access/CoreRoles.sol";
import {ChainlinkFeeds} from "./ChainlinkFeeds.sol";
import {IEverestOracle} from "../../src/interfaces/IEverestOracle.sol";

// How to add a new chain with its domain?
// 1. Add a domain ID. Domains originate from Connext (check their docs)
// 2. Create a fork and save it in "forks" mapping
// 3. Create a registry entry with addresses for reuiqred resources

contract ForkingSetup is CoreRoles, Test, ChainlinkFeeds {
  uint256 public constant ALICE_PK = 0xA;
  address public ALICE = vm.addr(ALICE_PK);
  uint256 public constant BOB_PK = 0xB;
  address public BOB = vm.addr(BOB_PK);
  uint256 public constant CHARLIE_PK = 0xC;
  address public CHARLIE = vm.addr(CHARLIE_PK);
  address public INITIALIZER = vm.addr(0x111A13); // arbitrary address

  uint32 originDomain;

  struct Registry {
    address weth;
    address usdc;
    address dai;
    address wmatic;
    address connext;
  }

  // domain => addresses registry
  mapping(uint256 => Registry) public registry;
  // domain => forkId
  mapping(uint256 => uint256) public forks;

  IVault public vault;
  Chief public chief;
  TimelockController public timelock;
  MockOracle mockOracle;
  IEverestOracle oracle;

  address public collateralAsset;
  address public debtAsset;

  uint256 initVaultShares;

  uint256 public constant DEFAULT_MAX_LTV = 75e16; // 75%
  uint256 public constant DEFAULT_LIQ_RATIO = 82.5e16; // 82.5%

  constructor() {
    vm.label(ALICE, "alice");
    vm.label(BOB, "bob");
    vm.label(CHARLIE, "charlie");
    vm.label(INITIALIZER, "initializer");

    forks[GOERLI_DOMAIN] = vm.createFork("goerli");
    forks[OPTIMISM_GOERLI_DOMAIN] = vm.createFork("optimism_goerli");
    forks[MUMBAI_DOMAIN] = vm.createFork("mumbai");
    forks[MAINNET_DOMAIN] = vm.createFork("mainnet");
    forks[OPTIMISM_DOMAIN] = vm.createFork("optimism");
    forks[ARBITRUM_DOMAIN] = vm.createFork("arbitrum");
    forks[POLYGON_DOMAIN] = vm.createFork("polygon");
    forks[GNOSIS_DOMAIN] = vm.createFork("gnosis");

    Registry memory goerli = Registry({
      weth: 0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6,
      usdc: 0xd35CCeEAD182dcee0F148EbaC9447DA2c4D449c4,
      dai: 0x2899a03ffDab5C90BADc5920b4f53B0884EB13cC,
      wmatic: address(0),
      connext: 0x99A784d082476E551E5fc918ce3d849f2b8e89B6
    });
    registry[GOERLI_DOMAIN] = goerli;

    Registry memory optimismGoerli = Registry({
      weth: 0x74c6FD7D2Bc6a8F0Ebd7D78321A95471b8C2B806,
      usdc: 0x3714A8C7824B22271550894f7555f0a672f97809,
      dai: 0xA8B4FBacE6B464f32daAf53b2b86dC91122194CB,
      wmatic: address(0),
      connext: 0x705791AD27229dd4CCf41b6720528AfE1bcC2910
    });
    registry[OPTIMISM_GOERLI_DOMAIN] = optimismGoerli;

    Registry memory mumbai = Registry({
      weth: 0xFD2AB41e083c75085807c4A65C0A14FDD93d55A9,
      usdc: 0xe6b8a5CF854791412c1f6EFC7CAf629f5Df1c747,
      dai: address(0),
      wmatic: address(0),
      connext: 0xfeBBcfe9a88aadefA6e305945F2d2011493B15b4
    });
    registry[MUMBAI_DOMAIN] = mumbai;

    Registry memory mainnet = Registry({
      weth: 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2,
      usdc: 0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48,
      dai: 0x6B175474E89094C44Da98b954EedeAC495271d0F,
      wmatic: address(0),
      connext: 0x8898B472C54c31894e3B9bb83cEA802a5d0e63C6
    });
    registry[MAINNET_DOMAIN] = mainnet;

    Registry memory optimism = Registry({
      weth: 0x4200000000000000000000000000000000000006,
      usdc: 0x7F5c764cBc14f9669B88837ca1490cCa17c31607,
      dai: address(0),
      wmatic: address(0),
      connext: 0x8f7492DE823025b4CfaAB1D34c58963F2af5DEDA
    });
    registry[OPTIMISM_DOMAIN] = optimism;

    Registry memory arbitrum = Registry({
      weth: 0x82aF49447D8a07e3bd95BD0d56f35241523fBab1,
      usdc: 0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8,
      wmatic: address(0),
      dai: address(0),
      connext: 0xEE9deC2712cCE65174B561151701Bf54b99C24C8
    });
    registry[ARBITRUM_DOMAIN] = arbitrum;

    Registry memory polygon = Registry({
      weth: 0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619,
      usdc: 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174,
      dai: 0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063,
      wmatic: 0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270,
      connext: 0x11984dc4465481512eb5b777E44061C158CF2259
    });
    registry[POLYGON_DOMAIN] = polygon;

    Registry memory gnosis = Registry({
      weth: 0x6A023CCd1ff6F2045C3309768eAd9E68F978f6e1,
      usdc: 0xDDAfbb505ad214D7b80b1f830fcCc89B60fb7A83,
      dai: 0xe91D153E0b41518A2Ce8Dd3D7944Fa863463a97d,
      wmatic: address(0),
      connext: 0x5bB83e95f63217CDa6aE3D181BA580Ef377D2109
    });
    registry[GNOSIS_DOMAIN] = gnosis;
  }

  function setUpFork(uint32 domain) public {
    Registry memory reg = registry[domain];
    if (reg.connext == address(0) && reg.weth == address(0) && reg.usdc == address(0)) {
      revert("No registry for this chain");
    }
    vm.selectFork(forks[domain]);

    originDomain = domain;

    oracle = new EverestOracle(
            assets[originDomain],
            priceFeeds[originDomain],
            address(chief)
        );

    if (reg.connext != address(0)) {
      vm.label(reg.connext, "Connext");
    }

    collateralAsset = reg.weth;
    vm.label(reg.weth, "WETH");

    if (reg.usdc == address(0)) {
      debtAsset = reg.dai;
      vm.label(debtAsset, "DAI");
    } else {
      debtAsset = reg.usdc;
      vm.label(debtAsset, "USDC");
    }

    if (domain == GOERLI_DOMAIN || domain == OPTIMISM_GOERLI_DOMAIN || domain == MUMBAI_DOMAIN) {
      MockERC20 tDAI = new MockERC20('Test DAI', 'tDAI');
      debtAsset = address(tDAI);
      vm.label(debtAsset, "testDAI");

      mockOracle = new MockOracle();

      // WETH assumed 2k usd/eth, dai = 1 usd
      mockOracle.setUSDPriceOf(collateralAsset, 2000e8);
      mockOracle.setUSDPriceOf(debtAsset, 1e8);
      oracle = mockOracle;
    }
  }

  function deploy(ILendingProvider[] memory providers) public {
    chief = new Chief(true, true);
    timelock = TimelockController(payable(chief.timelock()));
    // Grant this address all roles.
    _grantRoleChief(REBALANCER_ROLE, address(this));
    _grantRoleChief(LIQUIDATOR_ROLE, address(this));

    vault = new BorrowingVault(
            collateralAsset,
            debtAsset,
            address(oracle),
            address(chief),
            'Everest-V2 WETH-USDC Vault Shares',
            'fv2WETHUSDC',
            providers,
            DEFAULT_MAX_LTV,
            DEFAULT_LIQ_RATIO
        );

    bytes memory executionCall =
      abi.encodeWithSelector(chief.setVaultStatus.selector, address(vault), true);
    _callWithTimelock(address(chief), executionCall);

    initVaultShares = 10 ether;
    _initializeVault(address(vault), INITIALIZER, initVaultShares);
  }

  function deployVault(
    address collateralAsset_,
    address debtAsset_,
    string memory collateralAssetName,
    string memory debtAssetName,
    ILendingProvider[] memory providers
  )
    internal
  {
    collateralAsset = collateralAsset_;
    debtAsset = debtAsset_;
    vm.label(collateralAsset_, collateralAssetName);
    vm.label(debtAsset, debtAssetName);

    string memory nameVault =
      string.concat("Everest-V2 ", collateralAssetName, "-", debtAssetName, " Vault Shares");
    string memory symbolVault = string.concat("fv2", collateralAssetName, debtAssetName);

    chief = new Chief(true, true);
    timelock = TimelockController(payable(chief.timelock()));
    // Grant this address all roles.
    _grantRoleChief(REBALANCER_ROLE, address(this));
    _grantRoleChief(LIQUIDATOR_ROLE, address(this));

    vault = new BorrowingVault(
            collateralAsset,
            debtAsset,
            address(oracle),
            address(chief),
            nameVault,
            symbolVault,
            providers,
            DEFAULT_MAX_LTV,
            DEFAULT_LIQ_RATIO
        );

    bytes memory executionCall =
      abi.encodeWithSelector(chief.setVaultStatus.selector, address(vault), true);
    _callWithTimelock(address(chief), executionCall);

    uint256 minAmount = BorrowingVault(payable(address(vault))).minAmount();
    initVaultShares = 10 * minAmount;

    _initializeVault(address(vault), INITIALIZER, initVaultShares);
  }

  function _initializeVault(address vault_, address initializer, uint256 assets) internal {
    BorrowingVault bVault = BorrowingVault(payable(vault_));
    address collatAsset_ = bVault.asset();

    deal(collatAsset_, initializer, assets /*,true*/ );

    vm.startPrank(initializer);
    SafeERC20.safeIncreaseAllowance(IERC20(collatAsset_), vault_, assets);
    bVault.initializeVaultShares(assets);
    vm.stopPrank();
  }

  function _getMinCollateralAmount(
    BorrowingVault vault_,
    uint256 debt
  )
    internal
    view
    returns (uint256)
  {
    uint256 price = oracle.getPriceOf(vault_.debtAsset(), vault_.asset(), vault_.debtDecimals());
    uint256 minCollateralAmount =
      (debt * 1e18 * 10 ** vault_.decimals()) / (vault_.maxLtv() * price) + 1;
    if (minCollateralAmount < vault_.minAmount()) {
      minCollateralAmount = vault_.minAmount();
    }

    return minCollateralAmount;
  }

  function _callWithTimelock(address target, bytes memory callData) internal {
    timelock.schedule(target, 0, callData, 0x00, 0x00, 1.5 days);
    vm.warp(block.timestamp + 2 days);
    timelock.execute(target, 0, callData, 0x00, 0x00);
    rewind(2 days);
  }

  function _grantRoleChief(bytes32 role, address account) internal {
    bytes memory sendData = abi.encodeWithSelector(chief.grantRole.selector, role, account);
    _callWithTimelock(address(chief), sendData);
  }

  function _setVaultProviders(IVault v, ILendingProvider[] memory providers) internal {
    bytes memory callData = abi.encodeWithSelector(IVault.setProviders.selector, providers);
    _callWithTimelock(address(v), callData);
  }

  function _setActiveProvider(IVault v, ILendingProvider provider) internal {
    bytes memory callData = abi.encodeWithSelector(IVault.setActiveProvider.selector, provider);
    _callWithTimelock(address(v), callData);
  }

  function _getPermitBorrowArgs(
    LibSigUtils.Permit memory permit,
    uint256 ownerPKey,
    address vault_
  )
    internal
    returns (uint256 deadline, uint8 v, bytes32 r, bytes32 s)
  {
    bytes32 structHash = LibSigUtils.getStructHashBorrow(permit);
    bytes32 digest =
      LibSigUtils.getHashTypedDataV4Digest(IVaultPermissions(vault_).DOMAIN_SEPARATOR(), structHash);
    (v, r, s) = vm.sign(ownerPKey, digest);
    deadline = permit.deadline;
  }

  function _getPermitWithdrawArgs(
    LibSigUtils.Permit memory permit,
    uint256 ownerPKey,
    address vault_
  )
    internal
    returns (uint256 deadline, uint8 v, bytes32 r, bytes32 s)
  {
    bytes32 structHash = LibSigUtils.getStructHashWithdraw(permit);
    bytes32 digest =
      LibSigUtils.getHashTypedDataV4Digest(IVaultPermissions(vault_).DOMAIN_SEPARATOR(), structHash);
    (v, r, s) = vm.sign(ownerPKey, digest);
    deadline = permit.deadline;
  }

  /**
   * @dev this function was created to avoid stack to deep in `_getDepositAndBorrowCallData`.
   */
  function _buildPermitAsBytes(
    address owner,
    uint256 ownerPKey,
    address operator,
    address receiver,
    uint256 amount,
    uint256 plusNonce,
    address vault_,
    bytes32 actionArgsHash
  )
    internal
    returns (bytes memory arg)
  {
    LibSigUtils.Permit memory permit = LibSigUtils.buildPermitStruct(
      owner, operator, receiver, amount, plusNonce, vault_, actionArgsHash
    );

    (uint256 deadline, uint8 v, bytes32 r, bytes32 s) =
      _getPermitBorrowArgs(permit, ownerPKey, vault_);

    arg = abi.encode(vault_, owner, receiver, amount, deadline, v, r, s);
  }

  function _getDepositAndBorrowCallData(
    address owner,
    uint256 ownerPKey,
    uint256 amount,
    uint256 borrowAmount,
    address router,
    address vault_
  )
    internal
    returns (bytes memory callData)
  {
    IRouter.Action[] memory actions = new IRouter.Action[](3);
    actions[0] = IRouter.Action.Deposit;
    actions[1] = IRouter.Action.PermitBorrow;
    actions[2] = IRouter.Action.Borrow;

    bytes[] memory args = new bytes[](3);
    args[0] = abi.encode(vault_, amount, owner, router);
    args[1] = LibSigUtils.getZeroPermitEncodedArgs(vault_, owner, owner, borrowAmount);
    args[2] = abi.encode(vault_, borrowAmount, owner, owner);

    bytes32 actionArgsHash = LibSigUtils.getActionArgsHash(actions, args);

    // Replace permit action arguments, now with the signature values.
    args[1] = _buildPermitAsBytes(
      owner, // owner
      ownerPKey, // owner pkey to sign
      router, // operator
      owner, // receiver
      borrowAmount, // amount
      0, // plus nonce
      vault_, // vault
      actionArgsHash // actions args hash
    );

    callData = abi.encode(actions, args);
  }

  function _getDepositAndBorrow(
    address beneficiary,
    uint256 beneficiaryPrivateKey,
    uint256 amount,
    uint256 borrowAmount,
    address router,
    address vault_
  )
    internal
    returns (IRouter.Action[] memory, bytes[] memory)
  {
    IRouter.Action[] memory actions = new IRouter.Action[](3);
    actions[0] = IRouter.Action.Deposit;
    actions[1] = IRouter.Action.PermitBorrow;
    actions[2] = IRouter.Action.Borrow;

    bytes[] memory args = new bytes[](3);
    args[0] = abi.encode(vault_, amount, beneficiary, router);
    args[1] = LibSigUtils.getZeroPermitEncodedArgs(vault_, beneficiary, beneficiary, borrowAmount);
    args[2] = abi.encode(vault_, borrowAmount, beneficiary, beneficiary);

    bytes32 actionArgsHash = LibSigUtils.getActionArgsHash(actions, args);

    // Replace permit action arguments, now with the signature values.
    args[1] = _buildPermitAsBytes(
      beneficiary,
      beneficiaryPrivateKey,
      router,
      beneficiary,
      borrowAmount,
      0,
      vault_,
      actionArgsHash
    );

    return (actions, args);
  }
}
