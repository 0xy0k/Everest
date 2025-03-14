// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "forge-std/console.sol";
import {ScriptPlus} from "./ScriptPlus.s.sol";
import {AgaveGnosis} from "../src/providers/gnosis/AgaveGnosis.sol";

contract RunGnosis is ScriptPlus {
  AgaveGnosis agave;

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
    setOrDeployFlasherBalancer(false);
    setOrDeployRebalancer(false);

    agave = AgaveGnosis(getAddress("Agave_Gnosis"));
    /*agave = new AgaveGnosis();*/
    /*saveAddress("Agave_Gnosis", address(agave));*/

    if (chief.allowedVaultFactory(address(factory))) {
      deployBorrowingVaults();
      setBorrowingVaults();
    }

    if (chief.allowedVaultFactory(address(yieldFactory))) {
      deployYieldVaults();
    }

    /*upgradeBorrowingImpl(false);*/

    /*setVaultNewRating("BorrowingVault-WETHUSDC", 55);*/
    /*rebalanceVault("BorrowingVault-WETHUSDC", compound, aaveV3);*/

    // If setting all routers at once, call after deploying all chians
    /*setConnextReceivers();*/

    vm.stopBroadcast();
  }
}
