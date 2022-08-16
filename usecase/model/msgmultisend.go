package model

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
)

type MsgMultiSendParams struct {
	Inputs  []MsgMultiSendInput  `json:"inputs"`
	Outputs []MsgMultiSendOutput `json:"outputs"`
}

type MsgMultiSendInput struct {
	Address string     `json:"address"`
	Amount  coin.Coins `json:"amount"`
}

type MsgMultiSendOutput struct {
	Address string     `json:"address"`
	Amount  coin.Coins `json:"amount"`
}
