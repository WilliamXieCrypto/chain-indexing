package command

import (
	entity_event "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
)

type CreateBlockReward struct {
	blockHeight int64
	validator   string
	amount      coin.DecCoins
}

func NewCreateBlockReward(blockHeight int64, validator string, amount coin.DecCoins) *CreateBlockReward {
	return &CreateBlockReward{
		blockHeight,
		validator,
		amount,
	}
}

// Name returns name of command
func (*CreateBlockReward) Name() string {
	return "CreateBlockReward"
}

// Version returns version of command
func (*CreateBlockReward) Version() int {
	return 1
}

// Exec process the command data and return the event accordingly
func (cmd *CreateBlockReward) Exec() (entity_event.Event, error) {
	event := event.NewBlockRewarded(cmd.blockHeight, cmd.validator, cmd.amount)
	return event, nil
}
