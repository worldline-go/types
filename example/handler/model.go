package handler

import (
	"database/sql"
	"encoding/json"

	"github.com/shopspring/decimal"
	"github.com/worldline-go/types"
)

type Train struct {
	ID          int64                     `db:"id"          goqu:"skipinsert"`
	Details     types.Map                 `db:"details"     json:"details,omitempty"`
	Additionals types.RawJSON             `db:"additionals" json:"additionals,omitempty"`
	Price       sql.Null[json.Number]     `db:"price"`
	LastPrice   sql.Null[decimal.Decimal] `db:"last_price"`
}
