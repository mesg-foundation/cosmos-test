package serviceapp

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetService{}, "serviceapp/SetService", nil)
	cdc.RegisterConcrete(MsgGetService{}, "serviceapp/GetService", nil)
}
