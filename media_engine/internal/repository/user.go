package repository

import (
	"context"
	"errors"
	"rtc_media_engine/internal/models"
)

func (r *Repository) CreateUser(user *models.User, ctx context.Context) error {
	sqlQuery := "INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4)"
	_, err := r.Pool.Exec(ctx, sqlQuery, user.Email, user.PasswordHash, user.FirstName, user.LastName)
	if err != nil {
		return errors.New("Could not create user")
	}
	return nil
}

func (r *Repository) GetUserByEmail(user *models.User, ctx context.Context) error {
	sqlQuery := `SELECT id, first_name, last_name, created_at, updated_at FROM users WHERE email = $1`
	err := r.Pool.QueryRow(ctx, sqlQuery, user.Email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return errors.New("Could not find user")
	}

	return nil
}

func (r *Repository) GetUserPasswordHash(user *models.User, ctx context.Context) error {
	sqlQuery := `SELECT password_hash FROM users WHERE email = $1`
	row := r.Pool.QueryRow(ctx, sqlQuery, user.Email)

	err := row.Scan(&user.PasswordHash)
	if err != nil {
		return errors.New("User has not registered")
	}

	return nil
}
