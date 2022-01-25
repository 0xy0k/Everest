package types

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Aliases: []*TokenAlias{
			NewTokenAlias("KEX", "Tsuki", "", 6, []string{"ukex", "mkex"}),
		},
	}
}
