package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB initializes the global connection pool
func InitDB(connString string, ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	// Best practice: Explicitly configure pool limits
	config.MaxConns = 25
	config.MinConns = 2

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
		return nil
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
		return nil
	}

	fmt.Println("Database connection established")

	err = Migrate(ctx, pool)
	if err != nil {
		return nil
	}

	return pool
}
