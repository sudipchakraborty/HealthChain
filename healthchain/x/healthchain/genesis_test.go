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

		RecordList: []types.Record{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RecordCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HealthchainKeeper(t)
	healthchain.InitGenesis(ctx, *k, genesisState)
	got := healthchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RecordList, got.RecordList)
	require.Equal(t, genesisState.RecordCount, got.RecordCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
