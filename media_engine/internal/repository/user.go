package repository

import (
	"context"
	"errors"
	"rtc_media_engine/internal/models"
)

func (r *Repository) CreateUser(user *models.User, ctx context.Context) error {
	sqlQuery := "INSERT INTO users (users.email, users.password_hash) VALUES ($1, $2)"
	_, err := r.Pool.Exec(ctx, sqlQuery, user.Email, user.PasswordHash)
	if err != nil {
		return errors.New("Could not create user")
	}
	return nil
}

func (r *Repository) GetUserByEmail(user *models.User, ctx context.Context) error {
	sqlQuery := `SELECT users.id, users.first_name, users.last_name, users.created_at, users.updated_at
				FROM users
				WHERE users.email = $1`
	err := r.Pool.QueryRow(ctx, sqlQuery, user.Email).Scan(user.UserID, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return errors.New("Could not find user")
	}

	return nil
}

func (r *Repository) GetUserPasswordHash(user *models.User, ctx context.Context) error {
	sqlQuery := "SELECT users.password_hash FROM users WHERE users.email = $1"
	row := r.Pool.QueryRow(ctx, sqlQuery, user.Email)

	err := row.Scan(user.PasswordHash)
	if err != nil {
		return errors.New("User has not registered")
	}

	return nil
}
