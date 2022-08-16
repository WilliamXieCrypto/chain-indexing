package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCConnectionOpenTry struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgConnectionOpenTryParams
}

func NewCreateMsgIBCConnectionOpenTry(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgConnectionOpenTryParams,
) *CreateMsgIBCConnectionOpenTry {
	return &CreateMsgIBCConnectionOpenTry{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCConnectionOpenTry) Name() string {
	return "CreateMsgIBCConnectionOpenTry"
}

func (*CreateMsgIBCConnectionOpenTry) Version() int {
	return 1
}

func (cmd *CreateMsgIBCConnectionOpenTry) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCConnectionOpenTry(cmd.msgCommonParams, cmd.params)
	return event, nil
}
