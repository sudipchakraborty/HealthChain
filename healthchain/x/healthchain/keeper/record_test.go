package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "healthchain/testutil/keeper"
	"healthchain/testutil/nullify"
	"healthchain/x/healthchain/keeper"
	"healthchain/x/healthchain/types"
)

func createNRecord(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Record {
	items := make([]types.Record, n)
	for i := range items {
		items[i].Id = keeper.AppendRecord(ctx, items[i])
	}
	return items
}

func TestRecordGet(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRecord(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRecordRemove(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRecord(ctx, item.Id)
		_, found := keeper.GetRecord(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRecordGetAll(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRecord(ctx)),
	)
}

func TestRecordCount(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	items := createNRecord(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRecordCount(ctx))
}
