package handler

import (
	"net/http"
	"subscriptions-service/internal/ports"
)

type Handler struct {
	svc ports.Service
}

func NewHandler(svc ports.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetTotalCost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}
