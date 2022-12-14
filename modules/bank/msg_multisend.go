package bank

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	"github.com/kaifei-bianjie/common-parser/types"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type (
	DocMsgMultiSend struct {
		Inputs   []Item   `bson:"inputs"`
		Outputs  []Item   `bson:"outputs"`
		TempData []string `bson:"-"`
	}
	Item struct {
		Address string       `bson:"address"`
		Coins   []types.Coin `bson:"coins"`
	}
)

func (m *DocMsgMultiSend) GetType() string {
	return MsgTypeMultiSend
}

func (m *DocMsgMultiSend) BuildMsg(v interface{}) {
	msg := v.(*MsgMultiSend)
	for _, one := range msg.Inputs {
		m.Inputs = append(m.Inputs, Item{Address: one.Address, Coins: types.BuildDocCoins(one.Coins)})
		m.TempData = append(m.TempData, one.Address)
	}
	for _, one := range msg.Outputs {
		m.Outputs = append(m.Outputs, Item{Address: one.Address, Coins: types.BuildDocCoins(one.Coins)})
		m.TempData = append(m.TempData, one.Address)
	}

}

func (m *DocMsgMultiSend) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMultiSend)

	for _, one := range msg.Inputs {
		addrs = append(addrs, one.Address)
	}
	for _, one := range msg.Outputs {
		addrs = append(addrs, one.Address)
	}

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
