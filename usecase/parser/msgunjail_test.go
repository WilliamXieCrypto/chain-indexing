package parser_test

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/WilliamXieCrypto/chain-indexing/entity/command"
	command_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/WilliamXieCrypto/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgUnjail", func() {

		It("should parse Msg commands when there is slashing.MsgUnjail in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_UNJAIL_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(usecase_parser_test.TX_MSG_UNJAIL_BLOCK_RESULTS_RESP)
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgUnjail(
				event.MsgCommonParams{
					BlockHeight: int64(381001),
					TxHash:      "2D2075ECAF45DB84052B6A98E84C5E7FE158AB5157364FD5A934FAC96B77C6B1",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgUnjailParams{
					ValidatorAddr: "tcrocncl1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5m4uxzk",
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcrocncl1gs80n8fpc5mc3ywkgfy93l23tg0gdqj5m4uxzk"}))
		})
	})
})
