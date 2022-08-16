package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgRevoke struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgRevokeParams
}

func NewCreateMsgRevoke(
	msgCommonParams event.MsgCommonParams,
	params model.MsgRevokeParams,
) *CreateMsgRevoke {
	return &CreateMsgRevoke{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgRevoke) Name() string {
	return "CreateMsgRevoke"
}

func (*CreateMsgRevoke) Version() int {
	return 1
}

func (cmd *CreateMsgRevoke) Exec() (entity_event.Event, error) {
	event := event.NewMsgRevoke(cmd.msgCommonParams, cmd.params)
	return event, nil
}
