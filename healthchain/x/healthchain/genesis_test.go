package healthchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "healthchain/testutil/keeper"
	"healthchain/testutil/nullify"
	"healthchain/x/healthchain"
	"healthchain/x/healthchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HealthchainKeeper(t)
	healthchain.InitGenesis(ctx, *k, genesisState)
	got := healthchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
