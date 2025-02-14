package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    stablecointypes "github.com/jamescrystal/mychain/x/stablecoin/types"
)

// MintStablecoin implements the MsgMintStablecoin handler
func (k msgServer) MintStablecoin(goCtx context.Context, msg *types.MsgMintStablecoin) (*types.MsgMintStablecoinResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)

    // Create new coin ("ustable" is the stablecoin denomination)
    coin := sdk.NewCoin("ustable", sdk.NewInt(msg.Amount))

    // Mint coins to the module account
    if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin)); err != nil {
        return nil, sdkerrors.Wrapf(err, "failed to mint coins")
    }

    // Convert sender address
    sender, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        return nil, sdkerrors.Wrapf(err, "invalid sender address")
    }

    // Send minted coins to the user
    if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.NewCoins(coin)); err != nil {
        return nil, sdkerrors.Wrapf(err, "failed to send coins to account")
    }

    return &types.MsgMintStablecoinResponse{}, nil
}
