package handler

import (
	"database/sql"
	"encoding/json"

	"github.com/shopspring/decimal"
	"github.com/worldline-go/types"
)

type Train struct {
	ID           int64                 `db:"id"            goqu:"skipinsert"`
	Details      types.Map             `db:"details"       json:"details,omitempty"`
	Additionals  types.RawJSON         `db:"additionals"   json:"additionals,omitempty"`
	Price        sql.Null[json.Number] `db:"price"`
	LastPrice    decimal.NullDecimal   `db:"last_price"` // sql.Null[decimal.Decimal] is also usable
	Rate         sql.Null[string]      `db:"rate"`
	CustomNumber sql.Null[string]      `db:"custom_number"`
	Data         types.JSON[*Data]     `db:"data"`
	Slice        types.Slice[string]   `db:"slice"`
}

type Data struct {
	X int `json:"x"`
}
