package view

import (
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/projection/view"
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
)

type ValidatorBlockCommitmentsTotal struct {
	*view.Total
}

func NewValidatorBlockCommitmentsTotal(rdbHandle *rdb.Handle) *ValidatorBlockCommitmentsTotal {
	return &ValidatorBlockCommitmentsTotal{
		view.NewTotal(rdbHandle, "view_validator_block_commitments_total"),
	}
}
