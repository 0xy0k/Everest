package testslashing

import (
	"github.com/TsukiCore/tsuki/x/slashing/types"
)

// TestParams construct default slashing params for tests.
// Have to change these parameters for tests
// lest the tests take forever
func TestParams() types.Params {
	params := types.DefaultParams()
	params.DowntimeInactiveDuration = 60 * 60

	return params
}
