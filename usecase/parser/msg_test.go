package parser_test

import (
	"fmt"
	"strings"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/model/genesis"

	"github.com/WilliamXieCrypto/chain-indexing/infrastructure/tendermint"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

func mustParseBlockResp(rawResp string) (*model.Block, *model.RawBlock) {
	block, rawBlock, err := tendermint.ParseBlockResp(strings.NewReader(rawResp))
	if err != nil {
		panic(fmt.Sprintf("error parsing block response: %v", err))
	}

	return block, rawBlock
}

func mustParseBlockResultsResp(rawResp string) *model.BlockResults {
	blockResults, err := tendermint.ParseBlockResultsResp(strings.NewReader(rawResp))

	if err != nil {
		panic(fmt.Sprintf("error parsing block results response: %v", err))
	}

	return blockResults
}

func mustParseGenesisResp(rawResp string, strictParsing bool) *genesis.Genesis {
	genesis, err := tendermint.ParseGenesisResp(strings.NewReader(rawResp), strictParsing)

	if err != nil {
		panic(fmt.Sprintf("error parsing block genesis response: %v", err))
	}

	return genesis
}
