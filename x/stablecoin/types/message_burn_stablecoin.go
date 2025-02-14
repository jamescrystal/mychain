package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnStablecoin{}

func NewMsgBurnStablecoin(creator string, amount uint64) *MsgBurnStablecoin {
	return &MsgBurnStablecoin{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgBurnStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
