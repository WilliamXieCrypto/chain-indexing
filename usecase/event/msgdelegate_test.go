package event_test

import (
	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/WilliamXieCrypto/chain-indexing/usecase/coin"
	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgDelegate", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3"
			anyValidatorAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyAmount := coin.MustNewCoinFromString("basetcro", "123456")
			anyAutoClaimedRewards := coin.MustNewCoinFromString("basetcro", "789")
			anyParams := model.MsgDelegateParams{
				DelegatorAddress:   anyDelegatorAddress,
				ValidatorAddress:   anyValidatorAddress,
				Amount:             anyAmount,
				AutoClaimedRewards: anyAutoClaimedRewards,
			}
			event := event_usecase.NewMsgDelegate(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_DELEGATE_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgDelegate)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_DELEGATE_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
			Expect(typedEvent.AutoClaimedRewards).To(Equal(anyAutoClaimedRewards))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3"
			anyValidatorAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyAmount := coin.MustNewCoinFromString("basetcro", "123456")
			anyAutoClaimedRewards := coin.MustNewCoinFromString("basetcro", "789")
			anyParams := model.MsgDelegateParams{
				DelegatorAddress:   anyDelegatorAddress,
				ValidatorAddress:   anyValidatorAddress,
				Amount:             anyAmount,
				AutoClaimedRewards: anyAutoClaimedRewards,
			}
			event := event_usecase.NewMsgDelegate(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_DELEGATE_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgDelegate)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_DELEGATE_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
