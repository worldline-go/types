package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("not found")

type Handler struct {
	db *goqu.Database
}

func New(db *sqlx.DB) *Handler {
	return &Handler{
		db: goqu.New("postgres", db),
	}
}

func (h *Handler) GetTrain(ctx context.Context, id int64) (*Train, error) {
	train := &Train{}

	exist, err := h.db.From("train").
		Where(goqu.Ex{"id": id}).
		Executor().
		ScanStructContext(ctx, train)
	if err != nil {
		return nil, fmt.Errorf("failed to get train: %w", err)
	}

	if !exist {
		return nil, ErrNotFound
	}

	return train, nil
}

func (h *Handler) CreateTrain(ctx context.Context, train *Train) (int64, error) {
	var id int64
	_, err := h.db.Insert("train").
		Rows(train).
		Returning("id").
		Executor().
		ScanValContext(ctx, &id)
	if err != nil {
		return 0, fmt.Errorf("failed to create train: %w", err)
	}

	return id, nil
}

func (h *Handler) UpdateTrain(ctx context.Context, id int64, train *Train) error {
	_, err := h.db.Update("train").
		Set(train).
		Where(goqu.Ex{"id": id}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to update train: %w", err)
	}

	return nil
}
