package crisis

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/cosmosmod-parser/modules"
)

type DocMsgVerifyInvariant struct {
	Sender              string `bson:"sender"`
	InvariantModuleName string `bson:"invariant_module_name" yaml:"invariant_module_name"`
	InvariantRoute      string `bson:"invariant_route" yaml:"invariant_route"`
}

func (m *DocMsgVerifyInvariant) GetType() string {
	return MsgTypeVerifyInvariant
}

func (m *DocMsgVerifyInvariant) BuildMsg(v interface{}) {
	msg := v.(*MsgVerifyInvariant)
	m.Sender = msg.Sender
	m.InvariantModuleName = msg.InvariantModuleName
	m.InvariantRoute = msg.InvariantRoute

}

func (m *DocMsgVerifyInvariant) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgVerifyInvariant)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
