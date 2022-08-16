package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgWithdrawDelegatorReward struct {
	msgCommonParams event.MsgCommonParams

	params model.MsgWithdrawDelegatorRewardParams
}

func NewCreateMsgWithdrawDelegatorReward(
	msgCommonParams event.MsgCommonParams,
	params model.MsgWithdrawDelegatorRewardParams,
) *CreateMsgWithdrawDelegatorReward {
	return &CreateMsgWithdrawDelegatorReward{
		msgCommonParams,

		params,
	}
}

func (_ *CreateMsgWithdrawDelegatorReward) Name() string {
	return "CreateMsgWithdrawDelegatorReward"
}

func (_ *CreateMsgWithdrawDelegatorReward) Version() int {
	return 1
}

func (cmd *CreateMsgWithdrawDelegatorReward) Exec() (entity_event.Event, error) {
	event := event.NewMsgWithdrawDelegatorReward(
		cmd.msgCommonParams,
		cmd.params,
	)
	return event, nil
}
