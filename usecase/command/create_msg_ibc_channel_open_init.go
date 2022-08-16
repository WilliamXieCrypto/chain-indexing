package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelOpenInit struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelOpenInitParams
}

func NewCreateMsgIBCChannelOpenInit(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelOpenInitParams,
) *CreateMsgIBCChannelOpenInit {
	return &CreateMsgIBCChannelOpenInit{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelOpenInit) Name() string {
	return "CreateMsgIBCChannelOpenInit"
}

func (*CreateMsgIBCChannelOpenInit) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelOpenInit) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelOpenInit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
