package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

// RegisterCodec registers concrete message types
func RegisterCodec(cdc *codec.LegacyAmino) {
    cdc.RegisterConcrete(&MsgMintStablecoin{}, "stablecoin/MintStablecoin", nil)
    cdc.RegisterConcrete(&MsgBurnStablecoin{}, "stablecoin/BurnStablecoin", nil)
}

// RegisterInterfaces registers the module interfaces
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgMintStablecoin{},
        &MsgBurnStablecoin{},
    )
}
