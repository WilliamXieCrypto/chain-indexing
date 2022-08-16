package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgVote struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgVoteParams
}

func NewCreateMsgVote(
	msgCommonParams event.MsgCommonParams,
	params model.MsgVoteParams,
) *CreateMsgVote {
	return &CreateMsgVote{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgVote) Name() string {
	return "CreateMsgVote"
}

// Version returns version of command
func (*CreateMsgVote) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgVote) Exec() (entity_event.Event, error) {
	event := event.NewMsgVote(cmd.msgCommonParams, cmd.params)
	return event, nil
}
