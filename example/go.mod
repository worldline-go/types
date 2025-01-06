module github.com/worldline-go/types/example

go 1.22

replace github.com/worldline-go/types => ../

require (
	github.com/doug-martin/goqu/v9 v9.19.0
	github.com/jackc/pgx/v5 v5.7.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/rs/zerolog v1.33.0
	github.com/shopspring/decimal v1.4.0
	github.com/worldline-go/igmigrator/v2 v2.2.0
	github.com/worldline-go/initializer v0.5.0
	github.com/worldline-go/types v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rakunlabs/into v0.4.0 // indirect
	github.com/worldline-go/logz v0.5.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
