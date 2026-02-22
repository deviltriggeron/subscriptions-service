package handler

import "subscriptions-service/internal/ports"

type Handler struct {
	svc ports.Service
}

func NewHandler(svc ports.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
