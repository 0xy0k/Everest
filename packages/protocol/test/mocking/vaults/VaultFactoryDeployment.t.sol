// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/console.sol";
import {MockingSetup} from "../MockingSetup.sol";
import {IVault} from "../../../src/interfaces/IVault.sol";
import {ILendingProvider} from "../../../src/interfaces/ILendingProvider.sol";
import {BorrowingVault} from "../../../src/vaults/borrowing/BorrowingVault.sol";
import {BorrowingVaultFactory} from "../../../src/vaults/borrowing/BorrowingVaultFactory.sol";
import {YieldVault} from "../../../src/vaults/yields/YieldVault.sol";
import {YieldVaultFactory} from "../../../src/vaults/yields/YieldVaultFactory.sol";

contract VaultFactoryDeploymentUnitTests is MockingSetup {
  BorrowingVaultFactory public bVaultFactory;
  YieldVaultFactory public yVaultFactory;

  function setUp() public {
    // Deploy and set up the {BorrowingVaultFactory}.
    bVaultFactory = new BorrowingVaultFactory(address(chief));

    bytes memory callData =
      abi.encodeWithSelector(chief.allowVaultFactory.selector, address(bVaultFactory), true);
    _callWithTimelock(address(chief), callData);

    callData = abi.encodeWithSelector(
      bVaultFactory.setContractCode.selector, vm.getCode("BorrowingVault.sol:BorrowingVault")
    );
    _callWithTimelock(address(bVaultFactory), callData);

    // Deploy and set up the {YieldVaultFactory}.
    yVaultFactory = new YieldVaultFactory(address(chief));
    callData =
      abi.encodeWithSelector(chief.allowVaultFactory.selector, address(yVaultFactory), true);
    _callWithTimelock(address(chief), callData);
  }

  function do_borrowingVaultParamsCheck(
    BorrowingVault vault_,
    string memory name_,
    string memory symbol_
  )
    internal
  {
    // Check deployed vault params correspond.
    assertTrue(address(vault_) != address(0));
    assertTrue(vault_.asset() == collateralAsset);
    assertTrue(vault_.debtAsset() == debtAsset);
    // Casting requires since `oracle()` and `chief()` are not external methods of an {IVault}
    assertTrue(address(BorrowingVault(payable(address(vault_))).oracle()) == address(oracle));
    assertTrue(address(BorrowingVault(payable(address(vault_))).chief()) == address(chief));
    // Compare hashes since string comparison is not possible in Solidity.
    string memory name__ = vault_.name();
    string memory symbol__ = vault_.symbol();
    console.log("name__", name__);
    console.log("symbol__", symbol__);
    assertTrue(keccak256(abi.encodePacked(name__)) == keccak256(abi.encodePacked(name_)));
    assertTrue(keccak256(abi.encodePacked(symbol__)) == keccak256(abi.encodePacked(symbol_)));
    assertTrue(vault_.activeProvider() == mockProvider);
    assertTrue(vault_.maxLtv() == DEFAULT_MAX_LTV);
    assertTrue(vault_.liqRatio() == DEFAULT_LIQ_RATIO);
  }

  function do_yieldVaultParamsCheck(
    IVault vault_,
    string memory name_,
    string memory symbol_
  )
    internal
  {
    // Check deployed vault params correspond.
    assertTrue(address(vault_) != address(0));
    assertTrue(vault_.asset() == collateralAsset);
    assertTrue(vault_.debtAsset() == address(0));
    // Casting requires since `oracle()` and `chief()` are not external methods of an {IVault}
    assertTrue(address(BorrowingVault(payable(address(vault_))).chief()) == address(chief));
    // Compare hashes since string comparison is not possible in Solidity.
    string memory name__ = vault_.name();
    string memory symbol__ = vault_.symbol();
    console.log("name__", name__);
    console.log("symbol__", symbol__);
    assertTrue(keccak256(abi.encodePacked(name__)) == keccak256(abi.encodePacked(name_)));
    assertTrue(keccak256(abi.encodePacked(symbol__)) == keccak256(abi.encodePacked(symbol_)));
    assertTrue(vault_.activeProvider() == mockProvider);
  }

  function test_borrowingVaultDirectDeployment() public {
    // Check the default deployment in {MockingSetup} is functional.
    do_borrowingVaultParamsCheck(vault, "Everest-V2 tWETH-tDAI BorrowingVault", "fbvtWETHtDAI");
  }

  function test_yieldVaultDirectDeployment() public {
    ILendingProvider[] memory providers = new ILendingProvider[](1);
    providers[0] = mockProvider;

    YieldVault yvault = new YieldVault(
      collateralAsset,
      address(chief),
      "Everest-V2 tWETH YieldVault",
      "fyvtWETH",
      providers
    );
    do_yieldVaultParamsCheck(IVault(yvault), "Everest-V2 tWETH YieldVault", "fyvtWETH");
  }

  function test_borrowingVaultDeploymentWithFactory() public {
    ILendingProvider[] memory providers = new ILendingProvider[](1);
    providers[0] = mockProvider;

    address vaultAddr = chief.deployVault(
      address(bVaultFactory),
      abi.encode(
        address(collateralAsset),
        address(debtAsset),
        address(oracle),
        providers,
        DEFAULT_MAX_LTV,
        DEFAULT_LIQ_RATIO
      ),
      95
    );

    do_borrowingVaultParamsCheck(
      BorrowingVault(payable(vaultAddr)), "Everest-V2 tWETH-tDAI BorrowingVault", "fbvtWETHtDAI"
    );
  }

  function test_yieldVaultDeploymentWithFactory() public {
    ILendingProvider[] memory providers = new ILendingProvider[](1);
    providers[0] = mockProvider;

    address vaultAddr =
      chief.deployVault(address(yVaultFactory), abi.encode(address(collateralAsset), providers), 95);

    do_yieldVaultParamsCheck(IVault(vaultAddr), "Everest-V2 tWETH YieldVault", "fyvtWETH");
  }
}
