package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgSubmitSoftwareUpgradeProposal struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgSubmitSoftwareUpgradeProposalParams
}

func NewCreateMsgSubmitSoftwareUpgradeProposal(
	msgCommonParams event.MsgCommonParams,
	params model.MsgSubmitSoftwareUpgradeProposalParams,
) *CreateMsgSubmitSoftwareUpgradeProposal {
	return &CreateMsgSubmitSoftwareUpgradeProposal{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgSubmitSoftwareUpgradeProposal) Name() string {
	return "CreateMsgSubmitSoftwareUpgradeProposal"
}

// Version returns version of command
func (*CreateMsgSubmitSoftwareUpgradeProposal) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgSubmitSoftwareUpgradeProposal) Exec() (entity_event.Event, error) {
	event := event.NewMsgSubmitSoftwareUpgradeProposal(cmd.msgCommonParams, cmd.params)
	return event, nil
}
