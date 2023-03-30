package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
)

func (m *ProposalJoinDapp) ProposalType() string {
	return tsukitypes.ProposalTypeJoinDapp
}

func (m *ProposalJoinDapp) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *ProposalJoinDapp) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *ProposalJoinDapp) ValidateBasic() error {
	return nil
}

func (m *ProposalUpsertDapp) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertDapp
}

func (m *ProposalUpsertDapp) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *ProposalUpsertDapp) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *ProposalUpsertDapp) ValidateBasic() error {
	return nil
}
