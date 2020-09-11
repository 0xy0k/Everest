package config

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AccountAddressPrefix is the prefix for account address
	AccountAddressPrefix = "tsuki"
	// AccountPubKeyPrefix is the prefix for account public key
	AccountPubKeyPrefix = "tsukipub"
	// ValidatorAddressPrefix is the prefix for validator address
	ValidatorAddressPrefix = "tsukivaloper"
	// ValidatorPubKeyPrefix is the prefix for validator public key
	ValidatorPubKeyPrefix = "tsukivaloperpub"
	// ConsNodeAddressPrefix is the prefix for cons node address
	ConsNodeAddressPrefix = "tsukivalcons"
	// ConsNodePubKeyPrefix is the prefix for cons node public key
	ConsNodePubKeyPrefix = "tsukivalconspub"
)

// SetConfig is a function to set configuration for cosmos sdk
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
