package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "healthchain/testutil/keeper"
	"healthchain/testutil/nullify"
	"healthchain/x/healthchain/types"
)

func TestRecordQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRecord(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRecordRequest
		response *types.QueryGetRecordResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRecordRequest{Id: msgs[0].Id},
			response: &types.QueryGetRecordResponse{Record: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRecordRequest{Id: msgs[1].Id},
			response: &types.QueryGetRecordResponse{Record: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRecordRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Record(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRecordQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.HealthchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNRecord(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRecordRequest {
		return &types.QueryAllRecordRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RecordAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Record), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Record),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RecordAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Record), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Record),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RecordAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Record),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RecordAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
