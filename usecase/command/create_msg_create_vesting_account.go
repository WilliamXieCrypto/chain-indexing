package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgCreateVestingAccount struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgCreateVestingAccountParams
}

func NewCreateMsgCreateVestingAccount(
	msgCommonParams event.MsgCommonParams,
	params model.MsgCreateVestingAccountParams,
) *CreateMsgCreateVestingAccount {
	return &CreateMsgCreateVestingAccount{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgCreateVestingAccount) Name() string {
	return "CreateMsgCreateVestingAccount"
}

func (*CreateMsgCreateVestingAccount) Version() int {
	return 1
}

func (cmd *CreateMsgCreateVestingAccount) Exec() (entity_event.Event, error) {
	event := event.NewMsgCreateVestingAccount(cmd.msgCommonParams, cmd.params)
	return event, nil
}
