package types

import (
	"github.com/TsukiCore/tsuki/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// verify interface at compile time
var _ sdk.Msg = &MsgActivate{}

// NewMsgActivate creates a new MsgActivate instance
//nolint:interfacer
func NewMsgActivate(validatorAddr sdk.ValAddress) *MsgActivate {
	return &MsgActivate{
		ValidatorAddr: validatorAddr.String(),
	}
}

func (msg MsgActivate) Route() string { return RouterKey }
func (msg MsgActivate) Type() string  { return types.MsgTypeActivate }
func (msg MsgActivate) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddr)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr.Bytes()}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgActivate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgActivate) ValidateBasic() error {
	if msg.ValidatorAddr == "" {
		return ErrBadValidatorAddr
	}

	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgUnpause{}

// NewMsgUnpause creates a new MsgUnpause instance
//nolint:interfacer
func NewMsgUnpause(validatorAddr sdk.ValAddress) *MsgUnpause {
	return &MsgUnpause{
		ValidatorAddr: validatorAddr.String(),
	}
}

func (msg MsgUnpause) Route() string { return RouterKey }
func (msg MsgUnpause) Type() string  { return types.MsgTypeUnpause }
func (msg MsgUnpause) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddr)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr.Bytes()}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnpause) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgUnpause) ValidateBasic() error {
	if msg.ValidatorAddr == "" {
		return ErrBadValidatorAddr
	}

	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgPause{}

// NewMsgPause creates a new MsgPause instance
//nolint:interfacer
func NewMsgPause(validatorAddr sdk.ValAddress) *MsgPause {
	return &MsgPause{
		ValidatorAddr: validatorAddr.String(),
	}
}

func (msg MsgPause) Route() string { return RouterKey }
func (msg MsgPause) Type() string  { return types.MsgTypePause }
func (msg MsgPause) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddr)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr.Bytes()}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgPause) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgPause) ValidateBasic() error {
	if msg.ValidatorAddr == "" {
		return ErrBadValidatorAddr
	}

	return nil
}

// verify interface at compile time
var _ sdk.Msg = &MsgRefuteSlashingProposal{}

// NewMsgRefuteSlashingProposal creates a new MsgRefuteSlashingProposal instance
func NewMsgRefuteSlashingProposal(sender sdk.AccAddress, validatorAddr sdk.ValAddress, refutation string) *MsgRefuteSlashingProposal {
	return &MsgRefuteSlashingProposal{
		Sender:     sender.String(),
		Validator:  validatorAddr.String(),
		Refutation: refutation,
	}
}

func (msg MsgRefuteSlashingProposal) Route() string { return RouterKey }
func (msg MsgRefuteSlashingProposal) Type() string  { return types.MsgTypePause }
func (msg MsgRefuteSlashingProposal) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgRefuteSlashingProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgRefuteSlashingProposal) ValidateBasic() error {
	if msg.Validator == "" {
		return ErrBadValidatorAddr
	}

	return nil
}
