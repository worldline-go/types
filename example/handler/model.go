package handler

import (
	"database/sql"

	"github.com/shopspring/decimal"
	"github.com/worldline-go/types"
)

type Train struct {
	ID          int64                     `db:"id"          goqu:"skipinsert"`
	Details     types.Map                 `db:"details"     json:"details,omitempty"`
	Additionals types.RawJSON             `db:"additionals" json:"additionals,omitempty"`
	Price       sql.Null[decimal.Decimal] `db:"price"`
}
