// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title AaveV3Arbitrum
 *
 * @author Everest
 *
 * @notice This contract allows interaction with AaveV3.
 */

import {AaveV3Common, IV3Pool, AaveEModeHelper} from "../AaveV3Common.sol";

contract AaveV3Arbitrum is AaveV3Common {
  ///@inheritdoc AaveV3Common
  function _getPool() internal pure override returns (IV3Pool) {
    return IV3Pool(0x794a61358D6845594F94dc1DB02A252b5b4814aD);
  }

  ///@inheritdoc AaveV3Common
  function _getAaveEModeHelper() internal pure override returns (AaveEModeHelper) {
    return AaveEModeHelper(0x79Cb8Bc3035e5328498BE0d6CA93A6A56F27a2aE);
  }

  ///@inheritdoc AaveV3Common
  function providerName() public pure override returns (string memory) {
    return "Aave_V3_Arbitrum";
  }
}
