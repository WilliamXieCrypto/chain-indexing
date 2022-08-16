package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgDeposit struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgDepositParams
}

func NewCreateMsgDeposit(
	msgCommonParams event.MsgCommonParams,
	params model.MsgDepositParams,
) *CreateMsgDeposit {
	return &CreateMsgDeposit{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgDeposit) Name() string {
	return "CreateMsgDeposit"
}

// Version returns version of command
func (*CreateMsgDeposit) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgDeposit) Exec() (entity_event.Event, error) {
	event := event.NewMsgDeposit(cmd.msgCommonParams, cmd.params)
	return event, nil
}
