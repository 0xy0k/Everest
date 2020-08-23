package app

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	AccountAddressPrefix   = "tsuki"
	AccountPubKeyPrefix    = "tsukipub"
	ValidatorAddressPrefix = "tsukivaloper"
	ValidatorPubKeyPrefix  = "tsukivaloperpub"
	ConsNodeAddressPrefix  = "tsukivalcons"
	ConsNodePubKeyPrefix   = "tsukivalconspub"
)

func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
