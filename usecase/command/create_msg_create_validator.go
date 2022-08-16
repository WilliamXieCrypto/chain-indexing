package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgCreateValidator struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgCreateValidatorParams
}

func NewCreateMsgCreateValidator(msgCommonParams event.MsgCommonParams, params model.MsgCreateValidatorParams) *CreateMsgCreateValidator {
	return &CreateMsgCreateValidator{
		msgCommonParams,
		params,
	}
}

func (*CreateMsgCreateValidator) Name() string {
	return "CreateMsgCreateValidator"
}

func (*CreateMsgCreateValidator) Version() int {
	return 1
}

func (cmd *CreateMsgCreateValidator) Exec() (entity_event.Event, error) {

	event := event.NewMsgCreateValidator(cmd.msgCommonParams, cmd.params)
	return event, nil
}
