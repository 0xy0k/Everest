package app

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

// ExportAppStateAndValidators export the state of Tsuki for a genesis file
func (app *TsukiApp) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {
	return servertypes.ExportedApp{}, nil
}
