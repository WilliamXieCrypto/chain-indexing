package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelOpenAck struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelOpenAckParams
}

func NewCreateMsgIBCChannelOpenAck(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelOpenAckParams,
) *CreateMsgIBCChannelOpenAck {
	return &CreateMsgIBCChannelOpenAck{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelOpenAck) Name() string {
	return "CreateMsgIBCChannelOpenAck"
}

func (*CreateMsgIBCChannelOpenAck) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelOpenAck) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelOpenAck(cmd.msgCommonParams, cmd.params)
	return event, nil
}
