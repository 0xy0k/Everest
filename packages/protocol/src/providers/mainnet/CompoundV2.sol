// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title CompoundV2
 *
 * @author Everest
 *
 * @notice This contract allows interaction with CompoundV2.
 *
 * @dev The IAddrMapper needs to be properly configured for CompoundV2
 */

import {LibCompoundV2} from "../../libraries/LibCompoundV2.sol";
import {ILendingProvider} from "../../interfaces/ILendingProvider.sol";
import {IVault} from "../../interfaces/IVault.sol";
import {ICToken} from "../../interfaces/compoundV2/ICToken.sol";
import {ICERC20} from "../../interfaces/compoundV2/ICERC20.sol";
import {ICETH} from "../../interfaces/compoundV2/ICETH.sol";
import {IComptroller} from "../../interfaces/compoundV2/IComptroller.sol";
import {IAddrMapper} from "../../interfaces/IAddrMapper.sol";
import {IWETH9} from "../../abstracts/WETH9.sol";
import {Address} from "openzeppelin-contracts/contracts/utils/Address.sol";

contract CompoundV2 is ILendingProvider {
  using Address for address;

  /// @dev Custom errors
  error CompoundV2__deposit_failed(uint256 status);
  error CompoundV2__borrow_failed(uint256 status);
  error CompoundV2__withdraw_failed(uint256 status);
  error CompoundV2__payback_failed(uint256 status);

  /**
   * @notice Returns the {AddrMapper} contract applicable to this provider.
   */
  function getMapper() public pure returns (IAddrMapper) {
    return IAddrMapper(0xb2f151D71BD0396891A68C15BF608445a6aC835e);
  }

  /// @inheritdoc ILendingProvider
  function providerName() public pure override returns (string memory) {
    return "Compound_V2";
  }

  /**
   * @dev Approves vault's assets as collateral for Compound Protocol.
   *
   * @param asset address of the 'asset' to be approved as collateral.
   */
  function _enterCollatMarket(address asset) internal {
    IComptroller comptroller = IComptroller(0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B);

    address[] memory markets = new address[](1);
    markets[0] = asset;
    comptroller.enterMarkets(markets);
  }

  /**
   * @dev Returns true/false wether the given token is/isn't WETH.
   *
   * @param asset address of the token
   */
  function _isWETH(address asset) internal pure returns (bool) {
    return asset == 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
  }

  /// @inheritdoc ILendingProvider
  function approvedOperator(
    address keyAsset,
    address,
    address
  )
    external
    view
    returns (address operator)
  {
    operator = getMapper().getAddressMapping(providerName(), keyAsset);
  }

  /// @inheritdoc ILendingProvider
  function deposit(uint256 amount, IVault vault) external returns (bool success) {
    address asset = vault.asset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    _enterCollatMarket(cTokenAddr);

    if (_isWETH(asset)) {
      ICETH cETH = ICETH(cTokenAddr);
      // unwrap WETH to ETH
      IWETH9(asset).withdraw(amount);

      // cEth reverts if mint unsuccessful
      cETH.mint{value: amount}();
    } else {
      ICERC20 cToken = ICERC20(cTokenAddr);

      uint256 status = cToken.mint(amount);
      if (status != 0) {
        revert CompoundV2__deposit_failed(status);
      }
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function borrow(uint256 amount, IVault vault) external returns (bool success) {
    address asset = vault.debtAsset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    ICToken cToken = ICToken(cTokenAddr);

    uint256 status = cToken.borrow(amount);
    if (status != 0) {
      revert CompoundV2__borrow_failed(status);
    }

    // wrap ETH to WETH
    if (_isWETH(asset)) {
      IWETH9(asset).deposit{value: amount}();
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function withdraw(uint256 amount, IVault vault) external returns (bool success) {
    address asset = vault.asset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    ICToken cToken = ICToken(cTokenAddr);

    uint256 status = cToken.redeemUnderlying(amount);
    if (status != 0) {
      revert CompoundV2__withdraw_failed(status);
    }

    // wrap ETH to WETH
    if (_isWETH(asset)) {
      IWETH9(asset).deposit{value: amount}();
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function payback(uint256 amount, IVault vault) external returns (bool success) {
    address asset = vault.debtAsset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    if (_isWETH(asset)) {
      ICETH cETH = ICETH(cTokenAddr);
      // unwrap WETH to ETH
      IWETH9(asset).withdraw(amount);

      cETH.repayBorrow{value: amount}();
    } else {
      ICERC20 cToken = ICERC20(cTokenAddr);

      uint256 status = cToken.repayBorrow(amount);
      if (status != 0) {
        revert CompoundV2__payback_failed(status);
      }
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function getDepositRateFor(IVault vault) external view returns (uint256 rate) {
    address asset = vault.asset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    // Block rate transformed for common mantissa for Everest in ray (1e27)
    // Note: Compound uses base 1e18
    uint256 ratePerBlock = ICToken(cTokenAddr).supplyRatePerBlock() * 10 ** 9;

    // 2612670 is the approx. number of blocks per year with new PoS ethereum
    rate = ratePerBlock * 2612670;
  }

  /// @inheritdoc ILendingProvider
  function getBorrowRateFor(IVault vault) external view returns (uint256 rate) {
    address asset = vault.debtAsset();
    address cTokenAddr = getMapper().getAddressMapping(providerName(), asset);

    // Block rate transformed for common mantissa for Everest in ray (1e27)
    // Note: Compound uses base 1e18
    uint256 ratePerBlock = ICToken(cTokenAddr).borrowRatePerBlock() * 10 ** 9;

    // 2612670 is the approx. number of blocks per year with new PoS ethereum
    rate = ratePerBlock * 2612670;
  }

  /// @inheritdoc ILendingProvider
  function getDepositBalance(address user, IVault vault) external view returns (uint256 balance) {
    address asset = vault.asset();
    ICToken cToken = ICToken(getMapper().getAddressMapping(providerName(), asset));
    balance = LibCompoundV2.viewUnderlyingBalanceOf(cToken, user);
  }

  /// @inheritdoc ILendingProvider
  function getBorrowBalance(address user, IVault vault) external view returns (uint256 balance) {
    address asset = vault.debtAsset();
    ICToken cToken = ICToken(getMapper().getAddressMapping(providerName(), asset));
    balance = LibCompoundV2.viewBorrowingBalanceOf(cToken, user);
  }
}
