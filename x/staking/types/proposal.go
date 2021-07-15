package types

import (
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ProposalTypeUnjailValidator = "UnjailValidator"

func NewProposalUnjailValidator(proposer sdk.AccAddress, hash, reference string) *ProposalUnjailValidator {
	return &ProposalUnjailValidator{
		Proposer:  proposer,
		Hash:      "",
		Reference: "",
	}
}

func (m *ProposalUnjailValidator) ProposalType() string {
	return ProposalTypeUnjailValidator
}

func (m *ProposalUnjailValidator) ProposalPermission() types.PermValue {
	return types.PermCreateUnjailValidatorProposal
}

func (m *ProposalUnjailValidator) VotePermission() types.PermValue {
	return types.PermVoteUnjailValidatorProposal
}

// ValidateBasic returns basic validation
func (m *ProposalUnjailValidator) ValidateBasic() error {
	return nil
}
