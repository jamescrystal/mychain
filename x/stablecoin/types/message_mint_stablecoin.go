package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintStablecoin{}

func NewMsgMintStablecoin(creator string, amount uint64) *MsgMintStablecoin {
	return &MsgMintStablecoin{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMintStablecoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
