package model

import "github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

type CompleteBondingParams struct {
	Delegator string
	Validator string
	Amount    coin.Coins
}
