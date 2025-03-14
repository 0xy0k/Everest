// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.15;

/**
 * @title  MorphoCompound
 *
 * @author Everest
 *
 * @notice This contract allows interaction with Morpho.
 *
 * @dev The IAddrMapper needs to be properly configured for CompoundV3
 */

import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import {IVault} from "../../interfaces/IVault.sol";
import {ILendingProvider} from "../../interfaces/ILendingProvider.sol";
import {IMorpho} from "../../interfaces/morpho/IMorpho.sol";
import {ILens} from "../../interfaces/morpho/ILens.sol";
import {IAddrMapper} from "../../interfaces/IAddrMapper.sol";

contract MorphoCompound is ILendingProvider {
  address public constant LENS = 0x930f1b46e1D081Ec1524efD95752bE3eCe51EF67;
  address public constant MORPHO = 0x8888882f8f843896699869179fB6E4f7e3B58888;

  /**
   * @dev Returns {IMorpho} contract to be able to interact with the protocol
   */
  function _getMorpho() internal pure returns (IMorpho) {
    return IMorpho(MORPHO);
  }

  /**
   * @dev Returns the {IAddrMapper} on this chain.
   */
  function _getAddrMapper() internal pure returns (IAddrMapper) {
    return IAddrMapper(0xb2f151D71BD0396891A68C15BF608445a6aC835e);
  }

  /**
   * @dev Returns Compound's underlying cToken associated with the asset.
   *
   * @param asset address of the token to be used as collateral/debt.   *
   */
  function _getCToken(address asset) internal view returns (address cToken) {
    cToken = _getAddrMapper().getAddressMapping("Compound_V2", asset);
  }

  /// @inheritdoc ILendingProvider
  function providerName() public pure override returns (string memory) {
    return "Morpho_Compound_V2";
  }

  /// @inheritdoc ILendingProvider
  function approvedOperator(
    address,
    address,
    address
  )
    external
    pure
    override
    returns (address operator)
  {
    operator = MORPHO;
  }

  /// @inheritdoc ILendingProvider
  function deposit(uint256 amount, IVault vault) external override returns (bool success) {
    _getMorpho().supply(_getCToken(vault.asset()), address(vault), amount);
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function borrow(uint256 amount, IVault vault) external override returns (bool success) {
    _getMorpho().borrow(_getCToken(vault.debtAsset()), amount);
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function withdraw(uint256 amount, IVault vault) external override returns (bool success) {
    _getMorpho().withdraw(_getCToken(vault.asset()), amount);
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function payback(uint256 amount, IVault vault) external override returns (bool success) {
    _getMorpho().repay(_getCToken(vault.debtAsset()), address(vault), amount);
    success = true;
  }

  /// @inheritdoc ILendingProvider
  function getDepositRateFor(IVault vault) external view override returns (uint256 rate) {
    (uint256 ratePerBlock,,) = ILens(LENS).getAverageSupplyRatePerBlock(_getCToken(vault.asset()));
    //no. of blocks per year
    //convert from 1e18 to 1e27
    rate = ratePerBlock * 2612670 * 10 ** 9;
  }

  /// @inheritdoc ILendingProvider
  function getBorrowRateFor(IVault vault) external view override returns (uint256 rate) {
    (uint256 ratePerBlock,,) =
      ILens(LENS).getAverageBorrowRatePerBlock(_getCToken(vault.debtAsset()));
    //no. of blocks per year
    //convert from 1e18 to 1e27
    rate = ratePerBlock * 2612670 * 10 ** 9;
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
    (,, balance) = ILens(LENS).getCurrentSupplyBalanceInOf(_getCToken(vault.asset()), user);
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
    (,, balance) = ILens(LENS).getCurrentBorrowBalanceInOf(_getCToken(vault.debtAsset()), user);
  }
}
