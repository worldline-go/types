package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/worldline-go/igmigrator/v2"
)

var (
	ConnMaxLifetime = 15 * time.Minute
	MaxIdleConns    = 3
	MaxOpenConns    = 3
)

// Connect attempts to connect to database server.
func Connect(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "pgx", "postgres://postgres@localhost:5432/postgres")
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
