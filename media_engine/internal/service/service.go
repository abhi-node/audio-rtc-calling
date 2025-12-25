package service

import (
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/repository"
)

type Service struct {
	r         *repository.Repository
	jwtSecret []byte
}

func NewService(config config.Config) *Service {
	return &Service{
		r:         repository.NewRepository(config.CONN_STRING),
		jwtSecret: []byte(config.JWT_SECRET),
	}
}
