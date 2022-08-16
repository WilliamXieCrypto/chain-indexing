package event

import (
	"bytes"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"

	jsoniter "github.com/json-iterator/go"

	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/luci/go-render/render"
)

const BLOCK_REWARDED = "BlockRewarded"

type BlockRewarded struct {
	event_entity.Base

	Validator string        `json:"validator"`
	Amount    coin.DecCoins `json:"amount"`
}

func NewBlockRewarded(blockHeight int64, validator string, amount coin.DecCoins) *BlockRewarded {
	return &BlockRewarded{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        BLOCK_REWARDED,
			Version:     1,
			BlockHeight: blockHeight,
		}),

		validator,
		amount,
	}

}
func (event *BlockRewarded) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *BlockRewarded) String() string {
	return render.Render(event)
}

func DecodeBlockRewarded(encoded []byte) (event_entity.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *BlockRewarded
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
