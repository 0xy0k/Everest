package types

import sdk "github.com/TsukiCore/cosmos-sdk/types"

var _ sdk.Msg = MsgClaimValidator{}

type MsgClaimValidator struct {
	Moniker   string // 64 chars max
	Website   string // 64 chars max
	Social    string // 64 chars max
	Identity  string // 64 chars max
	Comission sdk.Dec
	ValKey    sdk.ValAddress
	PubKey    sdk.AccAddress
}

func NewMsgClaimValidator(
	moniker string,
	website string,
	social string,
	identity string,
	comission sdk.Dec,
	valKey sdk.ValAddress,
	pubKey sdk.AccAddress,
) *MsgClaimValidator {
	return &MsgClaimValidator{
		Moniker:   moniker,
		Website:   website,
		Social:    social,
		Identity:  identity,
		Comission: comission,
		ValKey:    valKey,
		PubKey:    pubKey,
	}
}

func (m MsgClaimValidator) Route() string {
	return ModuleName
}

func (m MsgClaimValidator) Type() string {
	return ClaimValidator
}

func (m MsgClaimValidator) ValidateBasic() error {
	panic("implement me")
}

func (m MsgClaimValidator) GetSignBytes() []byte {
	panic("implement me")
}

func (m MsgClaimValidator) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.PubKey,
	}
}

func (m MsgClaimValidator) Reset() {
	panic("implement me")
}

func (m MsgClaimValidator) String() string {
	panic("implement me")
}

func (m MsgClaimValidator) ProtoMessage() {
	panic("implement me")
}
