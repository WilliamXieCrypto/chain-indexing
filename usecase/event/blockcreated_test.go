package event_test

import (
	"encoding/json"

	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"

	random "github.com/brianvoe/gofakeit/v5"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

var _ = Describe("BlockCreated", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeBlockCreated", func() {
		It("should able to encode and decode to the same Event", func() {
			var block model.Block
			random.Struct(&block)
			event := event_usecase.NewBlockCreated(&block)
			encoded, _ := json.Marshal(event)

			decodedEvent, err := registry.DecodeByType(
				event_usecase.BLOCK_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
		})
	})
})
