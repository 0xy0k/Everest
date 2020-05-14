package createOrderBook

import (
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	errors "github.com/TsukiCore/cosmos-sdk/types/errors"
	"github.com/asaskevich/govalidator"

	"github.com/TsukiCore/tsuki/x/tsukiHub/constants"
)

type Message struct {
	Index int32            `json:"index" yaml:"index" valid:"required~index"`
	Base sdk.Coins		   `json:"base"  yaml:"base"  valid:"required~base"`
	Quote sdk.Coins		   `json:"quote" yaml:"quote" valid:"required~quote"`
	Mnemonic string 	   `json:"mnemonic" yaml:"mnemonic" valid:"required~mnemonic"`
	Curator sdk.AccAddress `json:"curator"  yaml:"curator" valid:"required~curator"`
}

var _ sdk.Msg = Message{}

func (message Message) Route() string { return constants.ModuleName }
func (message Message) Type() string  { return constants.createOrderBook }

func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessageCode, Error.Error())
	}
	return nil
}

func (message Message) GetSignBytes() []byte {
	return sdk.MustSortJSON(PackageCodec.MustMarshalJSON(message))
}

func (message Message) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{message.Curator}
}