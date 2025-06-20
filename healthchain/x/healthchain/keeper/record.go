package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"healthchain/x/healthchain/types"
)

// GetRecordCount get the total number of record
func (k Keeper) GetRecordCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RecordCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRecordCount set the total number of record
func (k Keeper) SetRecordCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RecordCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRecord appends a record in the store with a new id and update the count
func (k Keeper) AppendRecord(
	ctx sdk.Context,
	record types.Record,
) uint64 {
	// Create the record
	count := k.GetRecordCount(ctx)

	// Set the ID of the appended value
	record.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	appendedValue := k.cdc.MustMarshal(&record)
	store.Set(GetRecordIDBytes(record.Id), appendedValue)

	// Update record count
	k.SetRecordCount(ctx, count+1)

	return count
}

// SetRecord set a specific record in the store
func (k Keeper) SetRecord(ctx sdk.Context, record types.Record) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	b := k.cdc.MustMarshal(&record)
	store.Set(GetRecordIDBytes(record.Id), b)
}

// GetRecord returns a record from its id
func (k Keeper) GetRecord(ctx sdk.Context, id uint64) (val types.Record, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	b := store.Get(GetRecordIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRecord removes a record from the store
func (k Keeper) RemoveRecord(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	store.Delete(GetRecordIDBytes(id))
}

// GetAllRecord returns all record
func (k Keeper) GetAllRecord(ctx sdk.Context) (list []types.Record) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Record
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRecordIDBytes returns the byte representation of the ID
func GetRecordIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRecordIDFromBytes returns ID in uint64 format from a byte array
func GetRecordIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
