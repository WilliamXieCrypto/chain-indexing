package model

import "github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

type MsgFundCommunityPoolParams struct {
	Depositor string     `json:"depositor"`
	Amount    coin.Coins `json:"amount"`
}
