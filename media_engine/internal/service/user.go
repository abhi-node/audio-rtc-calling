package service

import (
	"context"
	"errors"
	"rtc_media_engine/internal/auth"
	"rtc_media_engine/internal/models"
)

func (s *Service) RegisterUser(user *models.User, rawPassword string, ctx context.Context) (string, error) {
	pass_hash, err := auth.HashPassword(rawPassword)
	if err != nil {
		return "", errors.New("Could not hash password")
	}

	user.PasswordHash = pass_hash
	err = s.r.CreateUser(user, ctx)
	if err != nil {
		return "", errors.New("Could not create user")
	}

	err = s.r.GetUserByEmail(user, ctx)
	if err != nil {
		return "", errors.New("Could not get user")
	}

	token, err := auth.GenerateToken(user.UserID, s.jwtSecret)
	if err != nil {
		return "", errors.New("Could not generate token")
	}

	return token, nil
}
