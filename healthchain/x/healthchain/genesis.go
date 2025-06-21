package healthchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"healthchain/x/healthchain/keeper"
	"healthchain/x/healthchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the record
	for _, elem := range genState.RecordList {
		k.SetRecord(ctx, elem)
	}

	// Set record count
	k.SetRecordCount(ctx, genState.RecordCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RecordList = k.GetAllRecord(ctx)
	genesis.RecordCount = k.GetRecordCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
