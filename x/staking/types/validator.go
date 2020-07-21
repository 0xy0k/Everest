package types

import sdk "github.com/TsukiCore/cosmos-sdk/types"

type Validator struct {
	Moniker  string
	Website  string
	Social   string
	Identity string

	Comission sdk.Dec

	ValKey sdk.ValAddress
	PubKey sdk.AccAddress
}
