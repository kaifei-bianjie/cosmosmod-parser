package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/kaifei-bianjie/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
