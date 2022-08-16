package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCChannelCloseInit struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgChannelCloseInitParams
}

func NewCreateMsgIBCChannelCloseInit(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgChannelCloseInitParams,
) *CreateMsgIBCChannelCloseInit {
	return &CreateMsgIBCChannelCloseInit{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCChannelCloseInit) Name() string {
	return "CreateMsgIBCChannelCloseInit"
}

func (*CreateMsgIBCChannelCloseInit) Version() int {
	return 1
}

func (cmd *CreateMsgIBCChannelCloseInit) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCChannelCloseInit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
