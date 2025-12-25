package handler

import (
	"context"
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(config config.Config, ctx context.Context) *Handler {
	return &Handler{
		s: service.NewService(config, ctx),
	}
}
