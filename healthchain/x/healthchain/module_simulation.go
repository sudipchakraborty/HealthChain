package healthchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"healthchain/testutil/sample"
	healthchainsimulation "healthchain/x/healthchain/simulation"
	"healthchain/x/healthchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = healthchainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateRecord = "op_weight_msg_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRecord int = 100

	opWeightMsgUpdateRecord = "op_weight_msg_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRecord int = 100

	opWeightMsgDeleteRecord = "op_weight_msg_record"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRecord int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	healthchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		RecordList: []types.Record{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		RecordCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&healthchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateRecord, &weightMsgCreateRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRecord = defaultWeightMsgCreateRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRecord,
		healthchainsimulation.SimulateMsgCreateRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateRecord, &weightMsgUpdateRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRecord = defaultWeightMsgUpdateRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRecord,
		healthchainsimulation.SimulateMsgUpdateRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRecord int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteRecord, &weightMsgDeleteRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRecord = defaultWeightMsgDeleteRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRecord,
		healthchainsimulation.SimulateMsgDeleteRecord(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRecord,
			defaultWeightMsgCreateRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				healthchainsimulation.SimulateMsgCreateRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRecord,
			defaultWeightMsgUpdateRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				healthchainsimulation.SimulateMsgUpdateRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRecord,
			defaultWeightMsgDeleteRecord,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				healthchainsimulation.SimulateMsgDeleteRecord(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
