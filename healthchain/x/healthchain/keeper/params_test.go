package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "healthchain/testutil/keeper"
	"healthchain/x/healthchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HealthchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
