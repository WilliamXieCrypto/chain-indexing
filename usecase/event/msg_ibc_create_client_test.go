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

	Describe("En/DecodeMsgIBCCreateClient", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2

			anySigner := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyParams := ibc_model.MsgCreateClientParams{
				MaybeTendermintLightClient: nil,

				Signer: anySigner,
			}

			event := event_usecase.NewMsgIBCCreateClient(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CREATE_CLIENT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCCreateClient)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CREATE_CLIENT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.Signer).To(Equal(anySigner))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2

			anySigner := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyParams := ibc_model.MsgCreateClientParams{
				MaybeTendermintLightClient: nil,

				Signer: anySigner,
			}

			event := event_usecase.NewMsgIBCCreateClient(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_IBC_CREATE_CLIENT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgIBCCreateClient)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_IBC_CREATE_CLIENT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Params.Signer).To(Equal(anySigner))
		})
	})
})
