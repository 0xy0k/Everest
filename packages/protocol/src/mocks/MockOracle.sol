// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import {IEverestOracle} from "../interfaces/IEverestOracle.sol";

contract MockOracle is IEverestOracle {
  mapping(address => mapping(address => uint256)) public prices;

  function getPriceOf(address currencyAsset, address commodityAsset, uint8 decimals)
    external
    view
    returns (uint256 price)
  {
    decimals;
    uint256 p = prices[currencyAsset][commodityAsset];
    price = p == 0 ? 1e18 : p;
  }

  function setPriceOf(address currencyAsset, address commodityAsset, uint256 price) public {
    prices[currencyAsset][commodityAsset] = price;
  }
}
