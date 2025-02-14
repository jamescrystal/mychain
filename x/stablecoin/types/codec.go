package types

import 

"github.com/cosmos/cosmos-sdk/codec"


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
