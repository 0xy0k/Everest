// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title DForceArbitrum
 *
 * @author Everest
 *
 * @notice This contract allows interaction with DForce.
 *
 * @dev The IAddrMapper needs to be properly configured for DForce.
 */

import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {IVault} from "../../interfaces/IVault.sol";
import {ILendingProvider} from "../../interfaces/ILendingProvider.sol";
import {IAddrMapper} from "../../interfaces/IAddrMapper.sol";
import {IComptroller} from "../../interfaces/compoundV2/IComptroller.sol";
import {IGenIToken} from "../../interfaces/dforce/IGenIToken.sol";
import {IIERC20} from "../../interfaces/dforce/IIERC20.sol";
import {IIETH} from "../../interfaces/dforce/IIETH.sol";
import {IWETH9} from "../../abstracts/WETH9.sol";
import {LibDForce} from "../../libraries/LibDForce.sol";

contract DForceArbitrum is ILendingProvider {
  /**
   * @dev Returns true/false wether the given token is/isn't WETH.
   *
   * @param token address of the 'token'
   */
  function _isWETH(address token) internal pure returns (bool) {
    return token == 0x82aF49447D8a07e3bd95BD0d56f35241523fBab1;
  }

  /**
   * @dev Returns the {IAddrMapper} on this chain.
   */
  function _getAddrmapper() internal pure returns (IAddrMapper) {
    return IAddrMapper(0x66211Ab72fB0a06e9E6eD8b21Aa3c1a01F171521);
  }

  /**
   * @dev Returns the Controller address of DForce.
   */
  function _getControllerAddress() internal pure returns (address) {
    return 0x8E7e9eA9023B81457Ae7E6D2a51b003D421E5408; // dForce Arbitrum
  }

  /**
   * @dev Approves vault's assets as collateral for dForce Protocol.
   *
   * @param _iTokenAddress address of the underlying {IGenIToken} to be approved as collateral
   */
  function _enterCollatMarket(address _iTokenAddress) internal {
    // Create a reference to the corresponding network Comptroller
    IComptroller controller = IComptroller(_getControllerAddress());

    address[] memory iTokenMarkets = new address[](1);
    iTokenMarkets[0] = _iTokenAddress;
    controller.enterMarkets(iTokenMarkets);
  }

  /**
   * @dev Returns DForce's underlying {IGenIToken} associated with the 'asset' to interact with DForce.
   *
   * @param asset address of the token to be used as collateral/debt
   */
  function _getiToken(address asset) internal view returns (address iToken) {
    iToken = _getAddrmapper().getAddressMapping(providerName(), asset);
  }

  /// @inheritdoc ILendingProvider
  function providerName() public pure override returns (string memory) {
    return "DForce_Arbitrum";
  }

  /// @inheritdoc ILendingProvider
  function approvedOperator(
    address keyAsset,
    address,
    address
  )
    external
    view
    override
    returns (address operator)
  {
    operator = _getiToken(keyAsset);
  }

  /// @inheritdoc ILendingProvider
  function deposit(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.asset();
    // Get iToken address from mapping
    address iTokenAddr = _getiToken(asset);

    // Enter and/or ensure collateral market is enacted
    _enterCollatMarket(iTokenAddr);

    if (_isWETH(asset)) {
      //unwrap WETH to ETH
      IWETH9(asset).withdraw(amount);

      // Create a reference to the iToken contract
      IIETH iToken = IIETH(iTokenAddr);

      // dForce protocol Mints iTokens, ETH method
      iToken.mint{value: amount}(address(this));
    } else {
      // Create a reference to the iToken contract
      IIERC20 iToken = IIERC20(iTokenAddr);

      // dForce Protocol mints iTokens
      iToken.mint(address(this), amount);
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function borrow(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.debtAsset();
    // Get iToken address from mapping
    address iTokenAddr = _getiToken(asset);

    // Create a reference to the corresponding iToken contract
    IGenIToken iToken = IGenIToken(iTokenAddr);

    // dForce Protocol Borrow Process, throw errow if not.
    iToken.borrow(amount);

    if (_isWETH(asset)) {
      // wrap ETH to WETH
      IWETH9(asset).deposit{value: amount}();
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function withdraw(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.asset();
    // Get iToken address from mapping
    address iTokenAddr = _getiToken(asset);

    // Create a reference to the corresponding iToken contract
    IGenIToken iToken = IGenIToken(iTokenAddr);

    // dForce Protocol Redeem Process, throw errow if not.
    iToken.redeemUnderlying(address(this), amount);

    if (_isWETH(asset)) {
      // wrap ETH to WETH
      IWETH9(asset).deposit{value: amount}();
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function payback(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.debtAsset();
    // Get iToken address from mapping
    address iTokenAddr = _getiToken(asset);

    if (_isWETH(asset)) {
      // Create a reference to the corresponding iToken contract
      IIETH iToken = IIETH(iTokenAddr);

      //unwrap WETH to ETH
      IWETH9(asset).withdraw(amount);

      iToken.repayBorrow{value: amount}();
    } else {
      // Create a reference to the corresponding iToken contract
      IIERC20 iToken = IIERC20(iTokenAddr);

      iToken.repayBorrow(amount);
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function getDepositRateFor(IVault vault) external view override returns (uint256 rate) {
    address iTokenAddr = _getiToken(vault.asset());

    // Block Rate transformed for common mantissa for Everest in ray (1e27), Note: dForce uses base 1e18
    uint256 bRateperBlock = IGenIToken(iTokenAddr).supplyRatePerBlock() * 10 ** 9;

    // The approximate number of blocks per year that is assumed by the dForce interest rate model
    uint256 blocksperYear = 2102400;
    rate = bRateperBlock * blocksperYear;
  }

  /// @inheritdoc ILendingProvider
  function getBorrowRateFor(IVault vault) external view override returns (uint256 rate) {
    address iTokenAddr = _getiToken(vault.debtAsset());

    // Block Rate transformed for common mantissa for Everest in ray (1e27), Note: dForce uses base 1e18
    uint256 bRateperBlock = IGenIToken(iTokenAddr).borrowRatePerBlock() * 10 ** 9;

    // The approximate number of blocks per year that is assumed by the dForce interest rate model
    // aligned with HundredFinance
    uint256 blocksperYear = 2336000;
    rate = bRateperBlock * blocksperYear;
  }

  /// @inheritdoc ILendingProvider
  function getDepositBalance(
    address user,
    IVault vault
  )
    external
    view
    override
    returns (uint256 balance)
  {
    address asset = vault.asset();
    IGenIToken iToken = IGenIToken(_getiToken(asset));
    balance = LibDForce.viewUnderlyingBalanceOf(iToken, user);
  }

  /// @inheritdoc ILendingProvider
  function getBorrowBalance(
    address user,
    IVault vault
  )
    external
    view
    override
    returns (uint256 balance)
  {
    address asset = vault.debtAsset();
    IGenIToken iToken = IGenIToken(_getiToken(asset));
    balance = LibDForce.viewBorrowingBalanceOf(iToken, user);
  }
}
