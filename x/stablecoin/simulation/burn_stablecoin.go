package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jamescrystal/mychain/x/stablecoin/keeper"
	"github.com/jamescrystal/mychain/x/stablecoin/types"
)

func SimulateMsgBurnStablecoin(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnStablecoin{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the BurnStablecoin simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "BurnStablecoin simulation not implemented"), nil, nil
	}
}
