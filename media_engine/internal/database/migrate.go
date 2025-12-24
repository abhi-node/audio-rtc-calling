package database

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(ctx context.Context, pool *pgxpool.Pool) error {

	migrationDir := "internal/database/migrations"
	migrations, err := os.ReadDir(migrationDir)
	if err != nil {
		return errors.New("Could not read directory")
	}

	for _, migration := range migrations {
		migrationPath := filepath.Join(migrationDir, migration.Name())
		sqlMigration, err := os.ReadFile(migrationPath)
		if err != nil {
			return errors.New("Could not read migration script")
		}

		_, err = pool.Exec(ctx, string(sqlMigration))
		if err != nil {
			return errors.New("Could not complete migration")
		}
	}

	return nil
}
