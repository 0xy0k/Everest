package types_test

import (
	"testing"

	"github.com/TsukiCore/tsuki/app"

	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgClaimValidator_ValidateBasic(t *testing.T) {
	app.SetConfig()
	valAddr1, err := types.ValAddressFromBech32("tsukivaloper15ky9du8a2wlstz6fpx3p4mqpjyrm5cgq38f2fp")
	require.NoError(t, err)

	pubKey, err := types.GetPubKeyFromBech32(types.Bech32PubKeyTypeConsPub, "tsukivalconspub1zcjduepqylc5k8r40azmw0xt7hjugr4mr5w2am7jw77ux5w6s8hpjxyrjjsq4xg7em")
	require.NoError(t, err)

	tests := []struct {
		name        string
		constructor func() (*stakingtypes.MsgClaimValidator, error)
	}{
		{
			name: "nil val key",
			constructor: func() (*stakingtypes.MsgClaimValidator, error) {
				return stakingtypes.NewMsgClaimValidator("me", types.NewDec(10), nil, pubKey)
			},
		},
		{
			name: "nil pub key",
			constructor: func() (*stakingtypes.MsgClaimValidator, error) {
				return stakingtypes.NewMsgClaimValidator("me", types.NewDec(10), valAddr1, nil)
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.constructor()
			require.Error(t, err)
		})
	}
}
