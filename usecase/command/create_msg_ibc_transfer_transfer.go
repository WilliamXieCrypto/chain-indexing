package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
)

type CreateMsgIBCTransferTransfer struct {
	msgCommonParams event.MsgCommonParams
	params          ibc_model.MsgTransferParams
}

func NewCreateMsgIBCTransferTransfer(
	msgCommonParams event.MsgCommonParams,
	params ibc_model.MsgTransferParams,
) *CreateMsgIBCTransferTransfer {
	return &CreateMsgIBCTransferTransfer{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgIBCTransferTransfer) Name() string {
	return "CreateMsgIBCTransferTransfer"
}

func (*CreateMsgIBCTransferTransfer) Version() int {
	return 1
}

func (cmd *CreateMsgIBCTransferTransfer) Exec() (entity_event.Event, error) {
	event := event.NewMsgIBCTransferTransfer(cmd.msgCommonParams, cmd.params)
	return event, nil
}
