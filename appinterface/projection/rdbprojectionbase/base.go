package rdbprojectionbase

import (
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
	"github.com/WilliamXieCrypto/chain-indexing/external/primptr"
	projection_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/projection"
)

const DEFAULT_TABLE = "projections"

// Table should have the following schema
// | Field                     | Data Type | Constraint  |
// | ------------------------- | --------- | ----------- |
// | id                        | VARCHAR   | PRIMARY KEY |
// | last_handled_event_height | INT64     | NOT NULL    |

// Base is a bas for projection which keeps track of last handled event height using relational
// database. It implements Id() and GetLastHandledEventHeight() of projection interface.
type Base struct {
	projection_usecase.Base

	rdbHandle *rdb.Handle
	store     *Store
}

// Create a new Base using table name in the RDb to keep the projection handling records
func NewRDbBase(rdbHandle *rdb.Handle, projectionId string) *Base {
	return NewRDbBaseWithOptions(
		rdbHandle,
		projectionId,
		Options{
			MaybeTable: primptr.String(DEFAULT_TABLE),
		},
	)
}

func NewRDbBaseWithOptions(rdbHandle *rdb.Handle, projectionId string, options Options) *Base {
	table := DEFAULT_TABLE
	if options.MaybeTable != nil {
		table = *options.MaybeTable
	}

	base := &Base{
		Base: projection_usecase.NewBase(
			projectionId,
		),

		rdbHandle: rdbHandle,
		store:     NewStore(table),
	}

	return base
}

type Options struct {
	// Customize table name in the RDb to keep the projection block handling records
	MaybeTable *string
}

func (base *Base) UpdateLastHandledEventHeight(rdbHandle *rdb.Handle, height int64) error {
	if err := base.store.UpdateLastHandledEventHeight(rdbHandle, base.Id(), height); err != nil {
		return err
	}
	return nil
}

// Implements projection.GetLastHandledEventHeight()
func (base *Base) GetLastHandledEventHeight() (*int64, error) {
	return base.store.GetLastHandledEventHeight(base.rdbHandle, base.Id())
}
