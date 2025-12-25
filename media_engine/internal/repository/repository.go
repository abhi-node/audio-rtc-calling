package repository

import (
	"context"
	"rtc_media_engine/internal/database"

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
