package tendermint

import (
	usecase_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model/genesis"
)

type Client interface {
	Genesis() (*genesis.Genesis, error)
	Block(height int64) (*usecase_model.Block, *usecase_model.RawBlock, error)
	BlockResults(height int64) (*usecase_model.BlockResults, error)
	LatestBlockHeight() (int64, error)
}
