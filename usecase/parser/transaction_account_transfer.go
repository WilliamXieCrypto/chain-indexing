package parser

import (
	"github.com/WilliamXieCrypto/chain-indexing/entity/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
	command_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser/utils"
)

func ParseTxAccountTransferCommands(
	blockHeight int64,
	txsResults []model.BlockResultsTxsResult,
) ([]command.Command, error) {
	commands := make([]command.Command, 0)
	for _, txsResult := range txsResults {
		var lastSender string
		for i, event := range txsResult.Events {
			if event.Type == "message" {
				messageEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])
				if messageEvent.HasAttribute("sender") {
					lastSender = messageEvent.MustGetAttributeByKey("sender")
				}
			} else if event.Type == "transfer" {
				transferEvent := utils.NewParsedTxsResultLogEvent(&txsResult.Events[i])

				amount := transferEvent.MustGetAttributeByKey("amount")
				if amount == "" {
					continue
				}

				var sender string
				if transferEvent.HasAttribute("sender") {
					sender = transferEvent.MustGetAttributeByKey("sender")
				} else {
					sender = lastSender
				}
				commands = append(commands, command_usecase.NewCreateAccountTransfer(
					blockHeight, model.AccountTransferParams{
						Recipient: transferEvent.MustGetAttributeByKey("recipient"),
						Sender:    sender,
						Amount:    coin.MustParseCoinsNormalized(amount),
					}))
			}
		}
	}

	return commands, nil
}
