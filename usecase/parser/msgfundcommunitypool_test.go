package parser_test

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/WilliamXieCrypto/chain-indexing/entity/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
	command_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/WilliamXieCrypto/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgFundCommunityPool", func() {
		It("should parse Msg commands when there is distribution.MsgFundCommunityPool in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_FUND_COMMUNITY_POOL_BLOCK_RESULTS_RESP,
			)
			accountAddressPrefix := "tcro"
			bondingDenom := "basetcro"

			pm := usecase_parser_test.InitParserManager()

			cmds, possibleSignerAddresses, err := parser.ParseBlockTxsMsgToCommands(
				pm,
				txDecoder,
				block,
				blockResults,
				accountAddressPrefix,
				bondingDenom,
			)
			Expect(err).To(BeNil())
			Expect(cmds).To(HaveLen(1))
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgFundCommunityPool(
				event.MsgCommonParams{
					BlockHeight: int64(460662),
					TxHash:      "933052FD68B10549312F3CBA9AF4F4CC77536BE5ECD335638DD36ECCE681201E",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgFundCommunityPoolParams{
					Depositor: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					Amount:    coin.MustParseCoinsNormalized("1basetcro"),
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"}))
		})
	})
})
