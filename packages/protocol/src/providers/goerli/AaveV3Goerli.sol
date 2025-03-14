// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title AaveV3Goerli
 *
 * @author Everest
 *
 * @notice This contract allows interaction with AaveV3.
 */

import {AaveV3Common, IV3Pool, AaveEModeHelper} from "../AaveV3Common.sol";

contract AaveV3Goerli is AaveV3Common {
  ///@inheritdoc AaveV3Common
  function _getPool() internal pure override returns (IV3Pool) {
    return IV3Pool(0x7b5C526B7F8dfdff278b4a3e045083FBA4028790);
  }

  ///@inheritdoc AaveV3Common
  function _getAaveEModeHelper() internal pure override returns (AaveEModeHelper) {
    return AaveEModeHelper(0xA169a53E5C8c36dD1CfeD0Ceb2BcD15026091272);
  }

  ///@inheritdoc AaveV3Common
  function providerName() public pure override returns (string memory) {
    return "Aave_V3_Goerli";
  }
}
