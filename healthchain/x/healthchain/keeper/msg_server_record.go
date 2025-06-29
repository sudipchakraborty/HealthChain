package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"healthchain/x/healthchain/types"
)

func (k msgServer) CreateRecord(goCtx context.Context, msg *types.MsgCreateRecord) (*types.MsgCreateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var record = types.Record{
		Creator:   msg.Creator,
		PatientId: msg.PatientId,
		Sys:       msg.Sys,
		Dia:       msg.Dia,
	}

	id := k.AppendRecord(
		ctx,
		record,
	)

	return &types.MsgCreateRecordResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateRecord(goCtx context.Context, msg *types.MsgUpdateRecord) (*types.MsgUpdateRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var record = types.Record{
		Creator:   msg.Creator,
		Id:        msg.Id,
		PatientId: msg.PatientId,
		Sys:       msg.Sys,
		Dia:       msg.Dia,
	}

	// Checks that the element exists
	val, found := k.GetRecord(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetRecord(ctx, record)

	return &types.MsgUpdateRecordResponse{}, nil
}

func (k msgServer) DeleteRecord(goCtx context.Context, msg *types.MsgDeleteRecord) (*types.MsgDeleteRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetRecord(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRecord(ctx, msg.Id)

	return &types.MsgDeleteRecordResponse{}, nil
}
