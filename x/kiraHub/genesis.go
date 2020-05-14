package tsukiHub

import sdkTypes "github.com/TsukiCore/cosmos-sdk/types"

type GenesisState struct{}

func DefaultGenesisState() GenesisState { return GenesisState{} }

func ValidateGenesis(genesisState GenesisState) error { return nil }

func InitializeGenesisState(context sdkTypes.Context, keeper Keeper, genesisState GenesisState) {
}
func ExportGenesis(context sdkTypes.Context, keeper Keeper) GenesisState {
	return GenesisState{}
}
