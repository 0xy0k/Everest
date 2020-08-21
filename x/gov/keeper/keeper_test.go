package keeper_test

import (
	"testing"

	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_SetPermissionsForRole(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perm := types.Permissions{
		Blacklist: nil,
		Whitelist: []uint32{
			types.PermClaimValidator,
		},
	}

	app.CustomGovKeeper.SetPermissionsForRole(ctx, types.RoleCouncilor, perm)
}
