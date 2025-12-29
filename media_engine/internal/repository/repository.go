package repository

import (
	"context"
	"errors"
	"rtc_media_engine/internal/database"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Pool *pgxpool.Pool
}

func NewRepository(connString string, ctx context.Context) *Repository {
	return &Repository{
		Pool: database.InitDB(connString, ctx),
	}
}

func (r *Repository) NewTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("Could not start transaction")
	}
	return tx, nil
}
