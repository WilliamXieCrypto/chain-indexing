package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgSubmitParamChangeProposal struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSubmitParamChangeProposalParams
}

func NewCreateMsgSubmitParamChangeProposal(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSubmitParamChangeProposalParams,
) *CreateMsgSubmitParamChangeProposal {
	return &CreateMsgSubmitParamChangeProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitParamChangeProposal) Name() string {
	return "CreateSubmitMsgParamChangeProposal"
}

// Version returns version of command
func (*CreateMsgSubmitParamChangeProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitParamChangeProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitParamChangeProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
