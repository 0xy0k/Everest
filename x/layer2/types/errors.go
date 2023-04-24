package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// layer2 module errors
var (
	ErrInvalidDappBondDenom                 = errors.Register(ModuleName, 1, "invalid dapp bond denom")
	ErrLowAmountToCreateDappProposal        = errors.Register(ModuleName, 2, "low amount to create dapp proposal")
	ErrDappDoesNotExist                     = errors.Register(ModuleName, 3, "dapp not found")
	ErrDappAlreadyExists                    = errors.Register(ModuleName, 4, "dapp already exists")
	ErrMaxDappBondReached                   = errors.Register(ModuleName, 5, "max dapp bond reached")
	ErrNotEnoughUserDappBond                = errors.Register(ModuleName, 6, "not enough user dapp bond")
	ErrUserDappBondDoesNotExist             = errors.Register(ModuleName, 7, "user dapp bond does not exist")
	ErrAlreadyADappVerifier                 = errors.Register(ModuleName, 8, "already a dapp verifier")
	ErrNotDappOperator                      = errors.Register(ModuleName, 9, "not a dapp operator")
	ErrAlreadyADappOperator                 = errors.Register(ModuleName, 10, "already a dapp operator")
	ErrDappNotHalted                        = errors.Register(ModuleName, 11, "dapp is not halted")
	ErrDappNotActive                        = errors.Register(ModuleName, 12, "dapp is not active")
	ErrDappNotPaused                        = errors.Register(ModuleName, 13, "dapp is not paused")
	ErrDappSessionDoesNotExist              = errors.Register(ModuleName, 14, "dapp session does not exist")
	ErrDappOperatorNotPaused                = errors.Register(ModuleName, 15, "dapp operator is not paused")
	ErrDappOperatorNotInActive              = errors.Register(ModuleName, 16, "dapp operator is not inactive")
	ErrDappOperatorNotActive                = errors.Register(ModuleName, 17, "dapp operator is not active")
	ErrNoDappSessionExists                  = errors.Register(ModuleName, 18, "dapp session does not exist")
	ErrNotDappSessionLeader                 = errors.Register(ModuleName, 19, "not a dapp session leader")
	ErrNumberOfOperatorsExceedsExecutorsMax = errors.Register(ModuleName, 20, "number of operators exceeds executors max")
	ErrOperatorJailed                       = errors.Register(ModuleName, 21, "operator jailed")
	ErrOperatorAlreadyExiting               = errors.Register(ModuleName, 22, "operator already exiting")
	ErrNotDappVerifier                      = errors.Register(ModuleName, 23, "not a dapp verifier")
	ErrLeaderCannotEvaluateSelfSubmission   = errors.Register(ModuleName, 24, "leader cannot evaluate self submission")
	ErrVerificationNotAllowedOnEmptySession = errors.Register(ModuleName, 25, "verification not allowed on empty session")
	ErrInvalidLpToken                       = errors.Register(ModuleName, 26, "invalid lp token")
	ErrOperationExceedsSlippage             = errors.Register(ModuleName, 27, "operation exceeds slippage")
)
