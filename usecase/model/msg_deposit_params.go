package model

import "github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

type MsgDepositParams struct {
	ProposalId string     `json:"proposalId"`
	Depositor  string     `json:"depositor"`
	Amount     coin.Coins `json:"amount"`
}
