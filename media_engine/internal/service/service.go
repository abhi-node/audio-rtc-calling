package service

import (
	"context"
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/repository"
)

type Service struct {
	r         *repository.Repository
	jwtSecret []byte
}

func NewService(config config.Config, ctx context.Context) *Service {
	return &Service{
		r:         repository.NewRepository(config.CONN_STRING, ctx),
		jwtSecret: []byte(config.JWT_SECRET),
	}
}
