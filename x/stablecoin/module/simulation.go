package stablecoin

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/jamescrystal/mychain/testutil/sample"
	stablecoinsimulation "github.com/jamescrystal/mychain/x/stablecoin/simulation"
	"github.com/jamescrystal/mychain/x/stablecoin/types"
)

// avoid unused import issue
var (
	_ = stablecoinsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgMintStablecoin = "op_weight_msg_mint_stablecoin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintStablecoin int = 100

	opWeightMsgBurnStablecoin = "op_weight_msg_burn_stablecoin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnStablecoin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	stablecoinGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&stablecoinGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintStablecoin int
	simState.AppParams.GetOrGenerate(opWeightMsgMintStablecoin, &weightMsgMintStablecoin, nil,
		func(_ *rand.Rand) {
			weightMsgMintStablecoin = defaultWeightMsgMintStablecoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintStablecoin,
		stablecoinsimulation.SimulateMsgMintStablecoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnStablecoin int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnStablecoin, &weightMsgBurnStablecoin, nil,
		func(_ *rand.Rand) {
			weightMsgBurnStablecoin = defaultWeightMsgBurnStablecoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnStablecoin,
		stablecoinsimulation.SimulateMsgBurnStablecoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintStablecoin,
			defaultWeightMsgMintStablecoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				stablecoinsimulation.SimulateMsgMintStablecoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBurnStablecoin,
			defaultWeightMsgBurnStablecoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				stablecoinsimulation.SimulateMsgBurnStablecoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
