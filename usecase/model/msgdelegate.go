package model

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
)

type MsgDelegateParams struct {
	DelegatorAddress   string    `json:"delegatorAddress"`
	ValidatorAddress   string    `json:"validatorAddress"`
	Amount             coin.Coin `json:"amount"`
	AutoClaimedRewards coin.Coin `json:"autoClaimedRewards"`
}
