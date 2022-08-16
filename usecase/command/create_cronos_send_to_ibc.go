package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

type CreateCronosSendToIBC struct {
	blockHeight int64
	params      model.CronosSendToIBCParams
}

func NewCreateCronosSendToIBC(
	blockHeight int64,
	params model.CronosSendToIBCParams,
) *CreateCronosSendToIBC {
	return &CreateCronosSendToIBC{
		blockHeight,
		params,
	}
}

// Name returns name of command
func (*CreateCronosSendToIBC) Name() string {
	return "CreateCronosSendToIBC"
}

// Version returns version of command
func (*CreateCronosSendToIBC) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateCronosSendToIBC) Exec() (entity_event.Event, error) {
	return event.NewCronosSendToIBCCreated(cmd.blockHeight, cmd.params), nil
}
