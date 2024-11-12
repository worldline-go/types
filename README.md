# types

Library for database types conversion.

```sh
go get github.com/worldline-go/types
```

## Usage

### types.Map

Map based on `map[string]interface{}` and null values stay as nil.

```go
type Train struct {
	Details types.Map `db:"details"`
}
```

---

### decimal.Decimal

To use `decimal.Decimal` type, you need to install `github.com/shopspring/decimal` package.

And need to add directly in driver, example for `pgx`, check in [connect.go](./example/database/connect.go).

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

Use with nullable `sql.Null` package or pointer.

```go
type Train struct {
	ID      int64                     `db:"id"      goqu:"skipinsert"`
	Details types.Map                 `db:"details"`
	Price   sql.Null[decimal.Decimal] `db:"price"`
}
```

## Development

Go to `example` folder and run `make` command and fallow usage.
