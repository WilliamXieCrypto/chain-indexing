package model

import "github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

type AccountTransferParams struct {
	Recipient string
	Sender    string
	Amount    coin.Coins
}
