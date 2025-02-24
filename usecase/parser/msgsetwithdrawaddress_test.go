package parser_test

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/WilliamXieCrypto/chain-indexing/entity/command"
	command_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/command"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/parser"
	usecase_parser_test "github.com/WilliamXieCrypto/chain-indexing/usecase/parser/test"
)

var _ = Describe("ParseMsgCommands", func() {
	Describe("MsgSetWithdrawAddress", func() {
		It("should parse Msg commands when there is distribution.MsgSetWithdrawAddress in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_SET_WITHDRAW_ADDRESS_BLOCK_RESULTS_RESP,
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
			Expect(cmds).To(Equal([]command.Command{command_usecase.NewCreateMsgSetWithdrawAddress(
				event.MsgCommonParams{
					BlockHeight: int64(460060),
					TxHash:      "9C2501310E18EE69A7FE5CA1A684A0701C43BEB1A8D91EDA80CC598C924F9CBE",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgSetWithdrawAddressParams{
					DelegatorAddress: "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
					WithdrawAddress:  "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh",
				},
			)}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"}))
		})
	})
})
