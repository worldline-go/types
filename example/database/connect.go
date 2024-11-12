package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/worldline-go/igmigrator/v2"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
)

var (
	ConnMaxLifetime = 15 * time.Minute
	MaxIdleConns    = 3
	MaxOpenConns    = 3
)

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

// Connect attempts to connect to database server.
func Connect(ctx context.Context) (*sqlx.DB, error) {
	dbDataSource := "postgres://postgres@localhost:5432/postgres"

	db, err := pgxConnect(ctx, dbDataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetConnMaxLifetime(ConnMaxLifetime)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	return db, nil
}

func MigrateDB(ctx context.Context, db *sqlx.DB) error {
	result, err := igmigrator.Migrate(ctx, db, &igmigrator.Config{})
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	for dir, r := range result.Path {
		if r.NewVersion != r.PrevVersion {
			log.Info().Msgf("path [%s] ran migrations from version %d to %d", dir, r.NewVersion, r.PrevVersion)
		}
	}

	return nil
}
