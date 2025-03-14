// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/console.sol";
import {ScriptPlus} from "./ScriptPlus.s.sol";
import {AaveV3Goerli} from "../src/providers/goerli/AaveV3Goerli.sol";

contract RunGoerli is ScriptPlus {
  AaveV3Goerli internal aaveV3;

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
    /*setOrDeployAddrMapper(false);*/
    /*setOrDeployFlasherBalancer(false);*/
    /*setOrDeployRebalancer(false);*/

    _setLendingProviders();

    if (chief.allowedVaultFactory(address(factory))) {
      deployBorrowingVaults();
      /*setBorrowingVaults();*/
    }

    /*upgradeBorrowingImpl(false);*/

    if (chief.allowedVaultFactory(address(yieldFactory))) {
      deployYieldVaults();
    }

    /*setVaultNewRating("BorrowingVault-WETHUSDC", 75);*/
    /*rebalanceVault("BorrowingVault-WETHUSDC", compound, aaveV3);*/

    // If setting all routers at once, call after deploying all chians
    /*setRouters();*/

    // setOrdeployAaveEModeHelper(false);

    vm.stopBroadcast();
  }

  function _setLendingProviders() internal {
    aaveV3 = AaveV3Goerli(getAddress("Aave_V3_Goerli"));
    /*aaveV3 = new AaveV3Goerli();*/
    /*saveAddress("Aave_V3_Goerli", address(aaveV3));*/
  }
}
