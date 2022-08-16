package view

import (
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/projection/view"
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
)

type BlockEventsTotal struct {
	*view.Total
}

func NewBlockEventsTotal(rdbHandle *rdb.Handle) *BlockEventsTotal {
	return &BlockEventsTotal{
		view.NewTotal(rdbHandle, "view_block_events_total"),
	}
}
