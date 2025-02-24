package event_test

import (
	"encoding/json"

	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeRawBlockCreated", func() {
		It("should able to encode and decode to the same Event", func() {
			var rawBlock model.RawBlock
			_ = json.Unmarshal([]byte(RAW_BLOCK_JSON), &rawBlock)
			event := event_usecase.NewEvent(&rawBlock)
			encoded, _ := json.Marshal(event)

			decodedEvent, err := registry.DecodeByType(
				event_usecase.RAW_BLOCK_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
		})
	})
})

const (
	RAW_BLOCK_JSON = `
{
  "block_id": {
    "hash": "82C25937191D1CF73BE9222CB04CE35B7A1366CC5BB08D9BB9AB457712E4F2D1",
    "parts": {
      "total": 1,
      "hash": "0A19FD939EBCE493C3126D99A6133F968AF454CAAEEAA1AD20EC4A3CFD60F2D5"
    }
  },
  "block": {
    "header": {
      "version": {
        "block": "11"
      },
      "chain_id": "testnet-croeseid-1",
      "height": "100",
      "time": "2020-10-15T09:33:42.195143319Z",
      "last_block_id": {
        "hash": "1532E4FFBDE4FE8CCDF5654A097D534B8C6E2EBC4473F36CFE314C0421970C2E",
        "parts": {
          "total": 1,
          "hash": "EEDFCCF098B1695CE939CF4E395AA8FC0EEC9F4673E1418B29E8928489BEF06A"
        }
      },
      "last_commit_hash": "E7437018C93A3BEA020C76DF06A209BFD9E2EEB6F01A5D4AF1EC64BA6488E6F5",
      "data_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
      "validators_hash": "C97D9876B52DAD0DA395EA0203EF1D85682198E82AB90D2CE60DB1FB25117C12",
      "next_validators_hash": "C97D9876B52DAD0DA395EA0203EF1D85682198E82AB90D2CE60DB1FB25117C12",
      "consensus_hash": "048091BC7DDC283F77BFBF91D73C44DA58C3DF8A9CBC867405D8B7F3DAADA22F",
      "app_hash": "6AE0920938F76727054BC2531247632C5C0521E2B91EA3A9864EA4FF55023D77",
      "last_results_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
      "evidence_hash": "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
      "proposer_address": "384E5F30F02538C0A34CBFF32F8D5554671C9029"
    },
    "data": {
      "txs": []
    },
    "evidence": {
      "evidence": []
    },
    "last_commit": {
      "height": "99",
      "round": 0,
      "block_id": {
        "hash": "1532E4FFBDE4FE8CCDF5654A097D534B8C6E2EBC4473F36CFE314C0421970C2E",
        "parts": {
          "total": 1,
          "hash": "EEDFCCF098B1695CE939CF4E395AA8FC0EEC9F4673E1418B29E8928489BEF06A"
        }
      },
      "signatures": [
        {
          "block_id_flag": 2,
          "validator_address": "384E5F30F02538C0A34CBFF32F8D5554671C9029",
          "timestamp": "2020-10-15T09:33:42.18646236Z",
          "signature": "Lnjz+jTGTk6DzqvbvdjFIG5zyzzpioiN37/B2pKi6ac/Kf2a+kmfeQA3CXnPO5GY/AoImfbVcCQs4asTSDCqCA=="
        },
        {
          "block_id_flag": 2,
          "validator_address": "4D9F47C5A19D5550685D2D55DF2FF8C5D7504CEB",
          "timestamp": "2020-10-15T09:33:42.195143319Z",
          "signature": "OpYxiF7QTAaG4NG2/iHWvts8ISQED6F+pNU5Cv2ts8c8mFW9J+g0oig3IXhvVG011uQjcr0CVw5P0eOXx1vYDg=="
        },
        {
          "block_id_flag": 2,
          "validator_address": "6D9E4B4995037D608E365CE90436C24580ABCC33",
          "timestamp": "2020-10-15T09:33:42.206633743Z",
          "signature": "jQd4JNrvX6DKmqDZ9VqoKtxRIxQHrvPWd4XW+ayrtVakiIMCWoVf1GMvxLbXYg68CyjmbuAX2VhCD0gSnj3pAw=="
        }
      ]
    }
  }
}`
)
