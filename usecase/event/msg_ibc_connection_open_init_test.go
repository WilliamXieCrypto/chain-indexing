package event_test

import (
	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	ibc_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model/ibc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgIBCConnectionOpenInit", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2

			anyParams := ibc_model.MsgConnectionOpenInitParams{
				RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
					ClientID: "07-tendermint-0",
					Counterparty: ibc_model.ConnectionCounterparty{
						ClientID:     "07-tendermint-0",
						ConnectionID: "",
						Prefix: ibc_model.ConnectionCounterpartyPrefix{
							KeyPrefix: []byte("aWJj"),
						},
					},
					Version: ibc_model.ConnectionCounterpartyVersion{
						Identifier: "1",
						Features: []string{
							"ORDER_ORDERED",
							"ORDER_UNORDERED",
						},
					},
					DelayPeriod: 0,
					Signer:      "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f",
				},

				ConnectionID: "connection-0",
			}

			event := event_usecase.NewMsgIBCConnectionOpenInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_INIT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_INIT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.Counterparty).To(Equal(anyParams.Counterparty))
			Expect(typedEvent.Params.Version).To(Equal(anyParams.Version))
			Expect(typedEvent.Params.DelayPeriod).To(Equal(anyParams.DelayPeriod))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2

			anyParams := ibc_model.MsgConnectionOpenInitParams{
				RawMsgConnectionOpenInit: ibc_model.RawMsgConnectionOpenInit{
					ClientID: "07-tendermint-0",
					Counterparty: ibc_model.ConnectionCounterparty{
						ClientID:     "07-tendermint-0",
						ConnectionID: "",
						Prefix: ibc_model.ConnectionCounterpartyPrefix{
							KeyPrefix: []byte("aWJj"),
						},
					},
					Version: ibc_model.ConnectionCounterpartyVersion{
						Identifier: "1",
						Features: []string{
							"ORDER_ORDERED",
							"ORDER_UNORDERED",
						},
					},
					DelayPeriod: 0,
					Signer:      "cro1gdswrmwtzgv3kvf28lvtt7qv7q7myzmn466r3f",
				},

				ConnectionID: "connection-0",
			}

			event := event_usecase.NewMsgIBCConnectionOpenInit(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CONNECTION_OPEN_INIT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCConnectionOpenInit)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CONNECTION_OPEN_INIT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.ClientID).To(Equal(anyParams.ClientID))
			Expect(typedEvent.Params.Counterparty).To(Equal(anyParams.Counterparty))
			Expect(typedEvent.Params.Version).To(Equal(anyParams.Version))
			Expect(typedEvent.Params.DelayPeriod).To(Equal(anyParams.DelayPeriod))
			Expect(typedEvent.Params.Signer).To(Equal(anyParams.Signer))
			Expect(typedEvent.Params.ConnectionID).To(Equal(anyParams.ConnectionID))
		})
	})
})
