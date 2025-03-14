// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title HundredGnosis
 *
 * @author Everest
 *
 * @notice This contract allows interaction with Hundred.
 *
 * @dev The IAddrMapper needs to be properly configured for DForce.
 */

import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {IVault} from "../../interfaces/IVault.sol";
import {ILendingProvider} from "../../interfaces/ILendingProvider.sol";
import {IAddrMapper} from "../../interfaces/IAddrMapper.sol";
import {IComptroller} from "../../interfaces/compoundV2/IComptroller.sol";
import {ICETH} from "../../interfaces/compoundV2/ICETH.sol";
import {ICERC20} from "../../interfaces/compoundV2/ICERC20.sol";
import {ICToken} from "../../interfaces/compoundV2/ICToken.sol";
import {ICETH} from "../../interfaces/compoundV2/ICETH.sol";
import {ICERC20} from "../../interfaces/compoundV2/ICERC20.sol";
import {IWETH9} from "../../abstracts/WETH9.sol";
import {LibCompoundV2} from "../../libraries/LibCompoundV2.sol";
import {IProxyReceiver} from "../../interfaces/compoundV2/IProxyReceiver.sol";

contract HundredGnosis is ILendingProvider {
  /// @dev Custom errors
  error Hundred__deposit_failed(uint256 status);
  error Hundred__payback_failed(uint256 status);
  error Hundred__withdraw_failed(uint256 status);
  error Hundred__borrow_failed(uint256 status);

  /**
   * @dev Returns true/false wether the given 'token' is/isn't WDAI.
   *
   * @param token address of the 'token'
   */
  function _isWDAI(address token) internal pure returns (bool) {
    return token == 0xe91D153E0b41518A2Ce8Dd3D7944Fa863463a97d;
  }

  /**
   * @dev Returns the {IAddrMapper} on this chain.
   */
  function _getAddrmapper() internal pure returns (IAddrMapper) {
    // TODO Define final address after deployment strategy is set.
    return IAddrMapper(0x2B3A214b2218f1ce35E94eE3BA14D8aF1712AB04);
  }

  /**
   * @dev Returns DForce's underlying {ICToken} associated with the 'asset' to interact with DForce.
   *
   * @param asset address of the token to be used as collateral/debt.
   */
  function _getCToken(address asset) internal view returns (address cToken) {
    cToken = _getAddrmapper().getAddressMapping("Hundred", asset);
  }

  /**
   * @dev Returns the Controller address of Hundred.
   */
  function _getComptrollerAddress() internal pure returns (address) {
    return 0x6bb6ebCf3aC808E26545d59EA60F27A202cE8586; // Hundred Gnosis
  }

  /**
   * @dev Returns the {ProxyReceiver}'s address used to withdraw from the protocol.
   */
  function _getProxyReceiver() internal pure returns (address) {
    return 0xcE04CdE2f1eB8177286F41479d753ab8B97322A9;
  }

  /**
   * @dev Approves vault's assets as collateral for Hundred Protocol.
   *
   * @param _cTokenAddress address of the asset to be approved as collateral.
   */
  function _enterCollatMarket(address _cTokenAddress) internal {
    // Create a reference to the corresponding network Comptroller
    IComptroller comptroller = IComptroller(_getComptrollerAddress());

    address[] memory cTokenMarkets = new address[](1);
    cTokenMarkets[0] = _cTokenAddress;
    comptroller.enterMarkets(cTokenMarkets);
  }

  /// @inheritdoc ILendingProvider
  function providerName() public pure override returns (string memory) {
    return "Hundred_Gnosis";
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
    operator = _getCToken(keyAsset);
  }

  /// @inheritdoc ILendingProvider
  function deposit(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.asset();
    address cTokenAddr = _getCToken(asset);

    _enterCollatMarket(cTokenAddr);

    if (_isWDAI(asset)) {
      ICETH cToken = ICETH(cTokenAddr);

      // unwrap WDAI to DAI
      IWETH9(asset).withdraw(amount);

      // Compound protocol Mints cTokens, ETH method
      cToken.mint{value: amount}();
    } else {
      ICERC20 cToken = ICERC20(cTokenAddr);

      uint256 status = cToken.mint(amount);
      if (status != 0) {
        revert Hundred__deposit_failed(status);
      }
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function borrow(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.debtAsset();
    address cTokenAddr = _getCToken(asset);

    ICToken cToken = ICToken(cTokenAddr);

    uint256 status = cToken.borrow(amount);

    if (status != 0) {
      revert Hundred__borrow_failed(status);
    }
    // wrap DAI to WDAI
    if (_isWDAI(asset)) {
      IWETH9(asset).deposit{value: amount}();
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function withdraw(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.asset();
    address cTokenAddr = _getCToken(asset);

    // Create a reference to the corresponding cToken contract
    ICToken cToken = ICToken(cTokenAddr);

    if (_isWDAI(asset)) {
      // use a proxy receiver because Hundred uses "to.transfer(amount)"
      // which runs out of gas with proxy contracts
      uint256 exRate = cToken.exchangeRateStored();
      uint256 scaledAmount = amount * 1e18 / exRate;

      address receiverAddr = _getProxyReceiver();
      cToken.transfer(receiverAddr, scaledAmount);
      IProxyReceiver(receiverAddr).withdraw(amount, cToken);
      // wrap DAI to WDAI
      IWETH9(asset).deposit{value: amount}();
    } else {
      uint256 status = cToken.redeemUnderlying(amount);
      if (status != 0) {
        revert Hundred__withdraw_failed(status);
      }
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function payback(uint256 amount, IVault vault) external override returns (bool success) {
    address asset = vault.debtAsset();
    address cTokenAddr = _getCToken(asset);

    if (_isWDAI(asset)) {
      ICETH cToken = ICETH(cTokenAddr);
      // unwrap WDAI to DAI
      IWETH9(asset).withdraw(amount);

      cToken.repayBorrow{value: amount}();
    } else {
      ICERC20 cToken = ICERC20(cTokenAddr);

      uint256 status = cToken.repayBorrow(amount);

      if (status != 0) {
        revert Hundred__payback_failed(status);
      }
    }
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function getDepositRateFor(IVault vault) external view override returns (uint256 rate) {
    address cTokenAddr = _getCToken(vault.asset());

    // Block Rate transformed for common mantissa for Everest in ray (1e27), Note: Compound uses base 1e18
    uint256 bRateperBlock = ICToken(cTokenAddr).supplyRatePerBlock() * 10 ** 9;

    // The approximate number of blocks per year that is assumed by the Compound interest rate model
    // https://docs.hundred.finance/developers/interest-rates/gnosis-chain
    uint256 blocksperYear = 6307200;
    rate = bRateperBlock * blocksperYear;
  }

  /// @inheritdoc ILendingProvider
  function getBorrowRateFor(IVault vault) external view override returns (uint256 rate) {
    address cTokenAddr = _getCToken(vault.debtAsset());

    // Block Rate transformed for common mantissa for Everest in ray (1e27), Note: Compound uses base 1e18
    uint256 bRateperBlock = ICToken(cTokenAddr).borrowRatePerBlock() * 10 ** 9;

    // The approximate number of blocks per year that is assumed by the Compound interest rate model
    // https://docs.hundred.finance/developers/interest-rates/gnosis-chain
    uint256 blocksperYear = 6307200;
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
    ICToken cToken = ICToken(_getCToken(asset));
    balance = LibCompoundV2.viewUnderlyingBalanceOf(cToken, user);
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
    ICToken cToken = ICToken(_getCToken(asset));
    balance = LibCompoundV2.viewBorrowingBalanceOf(cToken, user);
  }
}
