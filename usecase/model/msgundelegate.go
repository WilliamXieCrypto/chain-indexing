package model

import (
	"github.com/WilliamXieCrypto/chain-indexing/external/utctime"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
)

type MsgUndelegateParams struct {
	DelegatorAddress      string           `json:"delegatorAddress"`
	ValidatorAddress      string           `json:"validatorAddress"`
	Amount                coin.Coin        `json:"amount"`
	MaybeUnbondCompleteAt *utctime.UTCTime `json:"unbondCompleteAt"`
	AutoClaimedRewards    coin.Coin        `json:"autoClaimedRewards"`
}
