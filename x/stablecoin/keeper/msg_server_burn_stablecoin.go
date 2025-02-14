package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stablecointypes "github.com/jamescrystal/mychain/x/stablecoin/types"
)

// BurnStablecoin implements MsgBurnStablecoin
func (k msgServer) BurnStablecoin(goCtx context.Context, msg *stablecointypes.MsgBurnStablecoin) (*stablecointypes.MsgBurnStablecoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Define the stablecoin
	coin := sdk.NewCoin("ustable", sdk.NewInt(msg.Amount))

	// Convert sender address
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "invalid sender address")
	}

	// Transfer coins from user to module account
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, stablecointypes.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to transfer coins to module")
	}

	// Burn the coins
	if err := k.bankKeeper.BurnCoins(ctx, stablecointypes.ModuleName, sdk.NewCoins(coin)); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to burn coins")
	}

	return &stablecointypes.MsgBurnStablecoinResponse{}, nil
}
