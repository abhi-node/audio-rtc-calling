package handler

import (
	"rtc_media_engine/internal/config"
	"rtc_media_engine/internal/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(config config.Config) *Handler {
	return &Handler{
		s: service.NewService(config),
	}
}
