package parser_test

import (
	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
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
	Describe("MsgDeposit", func() {
		It("should parse gov.MsgDeposit command with effective height in the transaction", func() {
			txDecoder := utils.NewTxDecoder()
			block, _ := mustParseBlockResp(usecase_parser_test.TX_MSG_DEPOSIT_BLOCK_RESP)
			blockResults := mustParseBlockResultsResp(
				usecase_parser_test.TX_MSG_DEPOSIT_BLOCK_RESULTS_RESP,
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

			Expect(cmds).To(Equal([]command.Command{
				command_usecase.NewCreateMsgDeposit(
					event.MsgCommonParams{
						BlockHeight: int64(493659),
						TxHash:      "52D0E10D951DA5124E899A988ED6D014AF750B951688C6A59853A97E4AFE69B4",
						TxSuccess:   true,
						MsgIndex:    0,
					},
					model.MsgDepositParams{
						ProposalId: "9",
						Depositor:  "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn",
						Amount:     coin.MustParseCoinsNormalized("2basetcro"),
					},
				),
			}))
			Expect(possibleSignerAddresses).To(Equal([]string{"tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"}))

		})
	})
})
