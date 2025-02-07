package types

import "github.com/shopspring/decimal"

// Set default decimal.MarshalJSONWithoutQuotes to true.
//   - This will remove the quotes from the json output.
func init() {
	decimal.MarshalJSONWithoutQuotes = true
}
