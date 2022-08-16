package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCTimeoutOnClose struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgTimeoutOnCloseParams
}

func NewCreateMsgIBCTimeoutOnClose(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgTimeoutOnCloseParams,
) *CreateMsgIBCTimeoutOnClose {
	return &CreateMsgIBCTimeoutOnClose{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCTimeoutOnClose) Name() string {
	return "CreateMsgIBCTimeoutOnClose"
}

func (*CreateMsgIBCTimeoutOnClose) Version() int {
	return 1
}

func (cmd *CreateMsgIBCTimeoutOnClose) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCTimeoutOnClose(cmd.msgCommonParams, cmd.params)
	return event, nil
}

type CreateMsgAlreadyRelayedIBCTimeoutOnClose struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgTimeoutOnCloseParams
}

func NewCreateMsgAlreadyRelayedIBCTimeoutOnClose(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgTimeoutOnCloseParams,
) *CreateMsgAlreadyRelayedIBCTimeoutOnClose {
	return &CreateMsgAlreadyRelayedIBCTimeoutOnClose{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgAlreadyRelayedIBCTimeoutOnClose) Name() string {
	return "CreateMsgAlreadyRelayedIBCTimeoutOnClose"
}

func (*CreateMsgAlreadyRelayedIBCTimeoutOnClose) Version() int {
	return 1
}

func (cmd *CreateMsgAlreadyRelayedIBCTimeoutOnClose) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCTimeoutOnClose(cmd.msgCommonParams, cmd.params)
	return event, nil
}
