package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"healthchain/x/healthchain/types"
)

func (k Keeper) RecordAll(goCtx context.Context, req *types.QueryAllRecordRequest) (*types.QueryAllRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var records []types.Record
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	recordStore := prefix.NewStore(store, types.KeyPrefix(types.RecordKey))

	pageRes, err := query.Paginate(recordStore, req.Pagination, func(key []byte, value []byte) error {
		var record types.Record
		if err := k.cdc.Unmarshal(value, &record); err != nil {
			return err
		}

		records = append(records, record)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRecordResponse{Record: records, Pagination: pageRes}, nil
}

func (k Keeper) Record(goCtx context.Context, req *types.QueryGetRecordRequest) (*types.QueryGetRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	record, found := k.GetRecord(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRecordResponse{Record: record}, nil
}
