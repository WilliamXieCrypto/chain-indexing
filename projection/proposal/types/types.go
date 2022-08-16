package types

import "github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

type CommunityPoolSpendData struct {
	RecipientAddress string     `json:"recipient"`
	Amount           coin.Coins `json:"amount"`
}
