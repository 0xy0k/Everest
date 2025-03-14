// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.15;

/**
 * @title YieldVault
 *
 * @author Everest
 *
 * @notice Implementation vault that handles pooled single sided asset for
 * lending and strategies seeking yield.
 * All debt handling functions return zero or revert accordingly.
 * User state is kept at vaults via token-shares compliant to ERC4626.
 * This vault can aggregate protocols that implement yield strategies.
 */

import {
  IERC20,
  IERC20Metadata
} from "openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {IVault} from "../../interfaces/IVault.sol";
import {ILendingProvider} from "../../interfaces/ILendingProvider.sol";
import {SafeERC20} from "openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol";
import {BaseVault} from "../../abstracts/BaseVault.sol";

contract YieldVault is BaseVault {
  using SafeERC20 for IERC20Metadata;
  /// @dev Custom Errors

  error YieldVault__notApplicable();
  error YieldVault__rebalance_invalidProvider();

  /**
   * @notice Constructor of a new {YieldVault}.
   *
   * @param asset_ this vault will handle as main asset (collateral)
   * @param chief_ that deploys and controls this vault
   * @param name_ string of the token-shares handled in this vault
   * @param symbol_ string of the token-shares handled in this vault
   * @param providers_ array that will initialize this vault
   *
   * @dev Requirements:
   * - Must be initialized with a set of providers.
   * - Must set first provider in `providers_` array as `activeProvider`.
   *
   */
  constructor(
    address asset_,
    address chief_,
    string memory name_,
    string memory symbol_,
    ILendingProvider[] memory providers_
  )
    BaseVault(asset_, chief_, name_, symbol_)
  {
    _setProviders(providers_);
    _setActiveProvider(providers_[0]);

    _pauseForceAllActions();
  }

  receive() external payable {}

  /*//////////////////////////////////////////
      Asset management: overrides IERC4626
  //////////////////////////////////////////*/

  /// @inheritdoc BaseVault
  function maxWithdraw(address owner) public view override returns (uint256) {
    if (paused(VaultActions.Withdraw)) {
      return 0;
    }
    return convertToAssets(balanceOf(owner));
  }

  /// @inheritdoc BaseVault
  function maxRedeem(address owner) public view override returns (uint256) {
    if (paused(VaultActions.Withdraw)) {
      return 0;
    }
    return balanceOf(owner);
  }

  /*///////////////////////////////
      Debt management overrides 
  ///////////////////////////////*/

  /// @inheritdoc BaseVault
  function debtDecimals() public pure override returns (uint8) {}

  /// @inheritdoc BaseVault
  function debtAsset() public pure override returns (address) {}

  /// @inheritdoc BaseVault
  function balanceOfDebt(address) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function balanceOfDebtShares(address owner) public pure override returns (uint256 debtShares) {}

  /// @inheritdoc BaseVault
  function totalDebt() public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function convertDebtToShares(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function convertToDebt(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function maxBorrow(address) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function maxPayback(address) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function maxMintDebt(address) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function maxBurnDebt(address) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function previewBorrow(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function previewMintDebt(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function previewPayback(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function previewBurnDebt(uint256) public pure override returns (uint256) {}

  /// @inheritdoc BaseVault
  function borrow(uint256, address, address) public pure override returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function mintDebt(uint256, address, address) public pure override returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function payback(uint256, address) public pure override returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function burnDebt(uint256, address) public pure override returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /*///////////////////////
      Borrow allowances
  ///////////////////////*/

  /// @inheritdoc BaseVault
  function borrowAllowance(
    address,
    address,
    address
  )
    public
    view
    virtual
    override
    returns (uint256)
  {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function increaseBorrowAllowance(
    address,
    address,
    uint256
  )
    public
    virtual
    override
    returns (bool)
  {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function decreaseBorrowAllowance(
    address,
    address,
    uint256
  )
    public
    virtual
    override
    returns (bool)
  {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc BaseVault
  function permitBorrow(
    address,
    address,
    uint256,
    uint256,
    bytes32,
    uint8,
    bytes32,
    bytes32
  )
    public
    pure
    override
  {
    revert YieldVault__notApplicable();
  }

  /*/////////////////
      Rebalancing
  /////////////////*/

  /// @inheritdoc IVault
  function rebalance(
    uint256 assets,
    uint256 debt,
    ILendingProvider from,
    ILendingProvider to,
    uint256 fee,
    bool setToAsActiveProvider
  )
    external
    hasRole(msg.sender, REBALANCER_ROLE)
    returns (bool)
  {
    if (!_isValidProvider(address(from)) || !_isValidProvider(address(to))) {
      revert YieldVault__rebalance_invalidProvider();
    }

    if (debt != 0) {
      revert YieldVault__notApplicable();
    }

    _checkRebalanceFee(fee, assets);

    _executeProviderAction(assets, "withdraw", from);
    _executeProviderAction(assets, "deposit", to);

    if (setToAsActiveProvider) {
      _setActiveProvider(to);
    }

    emit VaultRebalance(assets, 0, address(from), address(to));
    return true;
  }

  /*////////////////////
       Liquidate
  ////////////////////*/

  /// @inheritdoc IVault
  function getHealthFactor(address) public pure returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc IVault
  function getLiquidationFactor(address) public pure returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /// @inheritdoc IVault
  function liquidate(address, address, uint256) public pure returns (uint256) {
    revert YieldVault__notApplicable();
  }

  /*/////////////////////////
      Admin set functions
  /////////////////////////*/

  /// @inheritdoc BaseVault
  function _setProviders(ILendingProvider[] memory providers) internal override {
    uint256 len = providers.length;
    for (uint256 i = 0; i < len;) {
      if (address(providers[i]) == address(0)) {
        revert BaseVault__setter_invalidInput();
      }
      _asset.forceApprove(
        providers[i].approvedOperator(asset(), asset(), debtAsset()), type(uint256).max
      );
      unchecked {
        ++i;
      }
    }
    _providers = providers;

    emit ProvidersChanged(providers);
  }
}
