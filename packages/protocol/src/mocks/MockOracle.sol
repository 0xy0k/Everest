// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

/**
 * @title MockOracle
 *
 * @author Fuijdao Labs
 *
 * @notice Mock implementation of the EverestOracle
 *
 * @dev Mock allows to set US prices openly for testing purposes.
 */

import {IEverestOracle} from "../interfaces/IEverestOracle.sol";

contract MockOracle is IEverestOracle {
  mapping(address => uint256) public prices;

  /// @inheritdoc IEverestOracle
  function getPriceOf(
    address currencyAsset,
    address commodityAsset,
    uint8 decimals
  )
    external
    view
    returns (uint256 price)
  {
    price = 10 ** uint256(decimals);

    if (commodityAsset != address(0)) {
      price = price * prices[commodityAsset];
    } else {
      price = price * (10 ** 8);
    }

    if (currencyAsset != address(0)) {
      price = price / prices[currencyAsset];
    } else {
      price = price / (10 ** 8);
    }
  }

  /**
   * @dev Set the USD price for testing environment.
   *
   * @param asset to which price will be set
   * @param price in USD for asset in 8 decimals
   */
  function setUSDPriceOf(address asset, uint256 price) public {
    prices[asset] = price;
  }
}
