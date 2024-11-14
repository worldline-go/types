# types

[![License](https://img.shields.io/github/license/worldline-go/types?color=red&style=flat-square)](https://raw.githubusercontent.com/worldline-go/types/main/LICENSE)
[![Coverage](https://img.shields.io/sonar/coverage/worldline-go_types?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=flat-square)](https://sonarcloud.io/summary/overall?id=worldline-go_types)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/worldline-go/types/test.yml?branch=main&logo=github&style=flat-square&label=ci)](https://github.com/worldline-go/types/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/worldline-go/types?style=flat-square)](https://goreportcard.com/report/github.com/worldline-go/types)
[![Go PKG](https://raw.githubusercontent.com/worldline-go/guide/main/badge/custom/reference.svg)](https://pkg.go.dev/github.com/worldline-go/types)

Library for types conversion.

```sh
go get github.com/worldline-go/types
```

## Usage

### types.Map

Map based on `map[string]interface{}` and null values stay as nil.

> This type not convert to base64 when marshaling, it is directly a json string.

```go
type Train struct {
	Details types.Map `db:"details"`
}
```

### types.RawJSON

`[]byte` type behind, same as `json.RawMessage` with scan and value methods.

```go
type Train struct {
	Details types.RawJSON `db:"details"`
}
```

---

### decimal.Decimal OR json.Number

To use `decimal.Decimal` type, you need to install `github.com/shopspring/decimal` package.  
Or use `json.Number` type for json number values and after that convert to `decimal.Decimal`.

Use with nullable `sql.Null` package or pointer.

And need to add directly in driver, example for `pgx`, check in [connect.go](./example/database/connect.go).

<details><summary>PGX Connect</summary>

```go
func pgxConnect(ctx context.Context, uri string) (*sqlx.DB, error) {
	connConfig, _ := pgx.ParseConfig(uri)
	afterConnect := stdlib.OptionAfterConnect(func(_ context.Context, conn *pgx.Conn) error {
		// Register decimal type
		pgxdecimal.Register(conn.TypeMap())

		return nil
	})

	db := sqlx.NewDb(stdlib.OpenDB(*connConfig, afterConnect), "pgx")

	err := db.PingContext(ctx)

	return db, err
}
```

In struct use like this:

```go

type Train struct {
	ID          int64                     `db:"id"          goqu:"skipinsert"`
	Details     types.Map                 `db:"details"     json:"details,omitempty"`
	Additionals types.RawJSON             `db:"additionals" json:"additionals,omitempty"`
	Price       sql.Null[json.Number]     `db:"price"`
	LastPrice   sql.Null[decimal.Decimal] `db:"last_price"`
}
```

</details>

## Development

Go to `example` folder and run `make` command and fallow usage.
