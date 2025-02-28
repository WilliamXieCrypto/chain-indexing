package block_test

import (
	"path"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	. "github.com/WilliamXieCrypto/chain-indexing/external/logger/test"
	"github.com/WilliamXieCrypto/chain-indexing/external/primptr"
	"github.com/WilliamXieCrypto/chain-indexing/external/utctime"
	"github.com/WilliamXieCrypto/chain-indexing/projection/block"
	block_view "github.com/WilliamXieCrypto/chain-indexing/projection/block/view"
	. "github.com/WilliamXieCrypto/chain-indexing/test"
	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
	usecase_model "github.com/WilliamXieCrypto/chain-indexing/usecase/model"
)

var BLOCK_MIGRATIONS_PATH = func() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("error retrieving file directory")
	}

	return path.Join(filename, "../migrations")
}()

var _ = Describe("Block", func() {
	WithProjectionTestEnv(func(testEnv ProjectionTestEnv) {
		pgxConn := testEnv.Conn

		reset := func() {
			// Only needs to run Reset() on one of the migrate instance because Reset() drops everything in the Database
			_ = testEnv.RootMigrate.Reset()
		}

		BeforeEach(func() {
			reset()

			testEnv.RootMigrate.MustUp()
			blockMigrate := testEnv.MigrateCreator(
				"block_schema_migrations",
				BLOCK_MIGRATIONS_PATH,
			)
			blockMigrate.MustUp()
		})

		AfterEach(func() {
			reset()
		})

		It("should project Blocks view when event is BlockCreated", func() {
			blocksView := block_view.NewBlocks(pgxConn.ToHandle())

			anyHeight := int64(1)
			event := event_usecase.NewBlockCreated(&usecase_model.Block{
				Height:          anyHeight,
				Hash:            "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:            utctime.FromUnixNano(int64(1000000)),
				AppHash:         "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				ProposerAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
				Txs: []string{
					"AAAMZqICtpjLrA3uEe3Rkg6cqDgQl0iBwG1Wm8ORZRzKL9EBAE1R0oP73H8AhHG1E1dke4ppXA/+43AyxWX7oVBTlQGnAABsiqVxCipneyzo9n1t7xauih3rXxssZTChLv/oPHUaSAAAAgAAAAAAAAAAAMw4kh8pZRDC442WIIUIo6+fIE1cTHkwEiv++LwFjNLwNtPTO1AVJx+dGivu9FZK3cN8r0mvLPVunExBxwKqKDkwVUbVSaIejuSc6Wm38rp+ekZjIExhf41FIwkocStPzBY4VYQ31YE7pmNnKocfCEg9CcKH+MIFYJ0FSxoEiMratLXjKAD6B/aMZ3pRCxxL/YxFNs+YKzGbgdKcLtydZ1B/OTGUeNU3SLNQJGmJeTL0AOwvQ4j7KkiQFZGa8uoQQEJo9rXjIg+xPk7a+zBdwC91EY/ZZz7pdVQMUgHBeKgsVvy0o78BOPWfuziBK5xVsoz5K8ZlMYQNPC5mY+bguMKYNJ5PQ9rzn+LseYT5jhalUsQrPABhEFOoQVUH1id3rszDkLLTn2/d89N/JJGY1+mL+upWWAsmJw1yTHqibSJ7RbiGffnh93MCOHeRe5OLrfmDFfhqOLNt9yuMYNdyJ+noVsfI7Ws9Kpxe2SnuCWBs9yOgM0l7UTMuIokqhGCkarXge/DUWLUcV694Jr8qcJT3OtoQohN+p+Xj66nowJLgFW7xBdFsT4vv7xI8giFxUpB+JkLgRZz2d2eam3iLDCHR+sNyvIDuXDUXhKM6aIykGgvHVHr3bBJRy/JPZFC0A3kqmheMnJpbV6kHXaIbmbgeQeXc+wxq1swFVxkwp14zbRmHwSGVGRgihjmoFYl9MyorQzNFETgp28gq2AKrCHNnIRs63m2cD5X4jZaFzfYAv9ifpOKqRtDgIFG0olge/ig00FFL6KdHySg1Qaab7g==",
				},
				Signatures: []usecase_model.BlockSignature{
					{
						BlockIdFlag:      2,
						ValidatorAddress: "F9E6FFB9B536956201AA138224FD888D03775AB4",
						Timestamp:        utctime.FromUnixNano(int64(1000000)),
						Signature:        "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
					},
					{
						BlockIdFlag:      2,
						ValidatorAddress: "031E3891DDB94FC7C7C132B7CD9736738110C889",
						Timestamp:        utctime.FromUnixNano(int64(2000000)),
						Signature:        "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
					},
				},
			})

			fakeLogger := NewFakeLogger()
			projection := block.NewBlock(fakeLogger, pgxConn, nil)

			Expect(blocksView.Count()).To(Equal(int64(0)))

			err := projection.HandleEvents(anyHeight, []event_entity.Event{event})
			Expect(err).To(BeNil())
			Expect(blocksView.Count()).To(Equal(int64(1)))

			actual, err := blocksView.FindBy(&block_view.BlockIdentity{
				MaybeHeight: primptr.Int64(anyHeight),
			})

			Expect(err).To(BeNil())
			Expect(actual).To(Equal(&block_view.Block{
				Height:           anyHeight,
				Hash:             "B69554A020537DA8E7C7610A318180C09BFEB91229BB85D4A78DDA2FACF68A48",
				Time:             utctime.FromUnixNano(int64(1000000)),
				AppHash:          "24474D86CBFA7E6328D473C17A9E46CD5A80FFE82A348A74844BF3E2BA2B3AF1",
				TransactionCount: 1,
				CommittedCouncilNodes: []block_view.BlockCommittedCouncilNode{
					{
						Address:    "F9E6FFB9B536956201AA138224FD888D03775AB4",
						Time:       utctime.FromUnixNano(int64(1000000)),
						Signature:  "ZW2pUcKFN/oPQCmdCouchXmgpPyd/Ddo45dhHEMwsBeHTBuSJh15zUMmfl5FZsPHeKC8citFvOm/52bgl5XHCw==",
						IsProposer: true,
					},
					{
						Address:    "031E3891DDB94FC7C7C132B7CD9736738110C889",
						Time:       utctime.FromUnixNano(int64(2000000)),
						Signature:  "uhWDC9NDT86FbRVGbOM2lGY8sVkWU51JJ9F8gPwTfK0ebcui1R34oM+jhPKdStn/4sq4qDgzbsN66cQ5kl8NAw==",
						IsProposer: false,
					},
				},
			}))
		})

		It("should update projection last handled event height when there is no event at the height", func() {
			anyHeight := int64(1)

			fakeLogger := NewFakeLogger()
			projection := block.NewBlock(fakeLogger, pgxConn, nil)

			Expect(projection.GetLastHandledEventHeight()).To(BeNil())

			err := projection.HandleEvents(anyHeight, []event_entity.Event{})
			Expect(err).To(BeNil())

			Expect(projection.GetLastHandledEventHeight()).To(Equal(primptr.Int64(anyHeight)))
		})
	})
})
