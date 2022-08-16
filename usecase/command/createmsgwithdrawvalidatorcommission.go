package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgWithdrawValidatorCommission struct {
	msgCommonParams event.MsgCommonParams

	params model.MsgWithdrawValidatorCommissionParams
}

func NewCreateMsgWithdrawValidatorCommission(
	msgCommonParams event.MsgCommonParams,
	params model.MsgWithdrawValidatorCommissionParams,
) *CreateMsgWithdrawValidatorCommission {
	return &CreateMsgWithdrawValidatorCommission{
		msgCommonParams,

		params,
	}
}

func (_ *CreateMsgWithdrawValidatorCommission) Name() string {
	return "CreateMsgWithdrawValidatorCommission"
}

func (_ *CreateMsgWithdrawValidatorCommission) Version() int {
	return 1
}

func (cmd *CreateMsgWithdrawValidatorCommission) Exec() (entity_event.Event, error) {
	event := event.NewMsgWithdrawValidatorCommission(
		cmd.msgCommonParams,
		cmd.params,
	)
	return event, nil
}
