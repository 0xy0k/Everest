package createOrderBook

import (
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/errors"
	"github.com/asaskevich/govalidator"

	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
)

type Message struct {
	Base string		   	   `json:"base"  yaml:"base"  valid:"required~Base is required"`
	Quote string		   `json:"quote" yaml:"quote" valid:"required~Quote is required"`
	Mnemonic string 	   `json:"mnemonic" yaml:"mnemonic" valid:"required~Mnemonic is required"`
	Curator string 		   `json:"curator"  yaml:"curator" valid:"required~Curator is required"`
}

var _ sdk.Msg = Message{}

func (message Message) Route() string { return constants.ModuleName }
func (message Message) Type() string  { return constants.CreateOrderBookTransaction }

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
	var curator, _ = sdk.AccAddressFromBech32(message.Curator)
	return []sdk.AccAddress{curator}
}