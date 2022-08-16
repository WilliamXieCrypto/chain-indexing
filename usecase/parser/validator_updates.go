package parser

import (
	"github.com/WilliamXieCrypto/chain-indexing/entity/command"
	command_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

func ParseValidatorUpdatesCommands(
	blockHeight int64,
	validatorUpdates []model.BlockResultsValidatorUpdate,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)

	for _, update := range validatorUpdates {
		power := "0"
		if update.MaybePower != nil {
			power = update.MaybePower.String()
		}
		commands = append(commands, command_usecase.NewChangePower(blockHeight, model.PowerChangeParams{
			TendermintPubkey: update.Pubkey.Pubkey,
			Power:            power,
		}))
	}

	return commands, nil
}
