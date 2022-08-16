package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateMsgNFTBurnNFT struct {
	msgCommonParams event.MsgCommonParams
	params          model.MsgNFTBurnNFTParams
}

func NewCreateMsgNFTBurnNFT(
	msgCommonParams event.MsgCommonParams,
	params model.MsgNFTBurnNFTParams,
) *CreateMsgNFTBurnNFT {
	return &CreateMsgNFTBurnNFT{
		msgCommonParams,
		params,
	}
}

// Name returns name of command
func (*CreateMsgNFTBurnNFT) Name() string {
	return "CreateMsgNFTBurnNFT"
}

// Version returns version of command
func (*CreateMsgNFTBurnNFT) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateMsgNFTBurnNFT) Exec() (entity_event.Event, error) {
	event := event.NewMsgNFTBurnNFT(cmd.msgCommonParams, cmd.params)
	return event, nil
}
