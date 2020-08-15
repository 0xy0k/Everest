package types

import sdk "github.com/TsukiCore/cosmos-sdk/types"

const (
	// ModuleName is the name of the custom staking
	ModuleName = "customstaking"

	ClaimValidator = "claim-validator"
)

var (
	ValidatorsKey          = []byte{0x21} // Validators key prefix.
	ValidatorsByMonikerKey = []byte{0x22} // Validators by moniker prefix.
)

// GetValidatorKey gets the key for the validator with address
func GetValidatorKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, operatorAddr.Bytes()...)
}

func GetValidatorByMonikerKey(moniker string) []byte {
	return append(ValidatorsByMonikerKey, []byte(moniker)...)
}
