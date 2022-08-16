package view

import (
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/projection/view"
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
)

type IBCChannelMessagesTotal interface {
	Increment(identity string, total int64) error
	SumBy(identities []string) (int64, error)
}

type IBCChannelMessagesTotalView struct {
	*view.Total
}

func NewIBCChannelMessagesTotalView(rdbHandle *rdb.Handle) IBCChannelMessagesTotal {
	return &IBCChannelMessagesTotalView{
		view.NewTotal(rdbHandle, "view_ibc_channel_messages_total"),
	}
}
