package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"subscriptions-service/internal/domain"
	"subscriptions-service/internal/dto"
	"subscriptions-service/internal/mapper"
	"subscriptions-service/internal/ports"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	var req dto.CreateSubscriptionReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sub, err := h.svc.Create(r.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInvalidPrice) || errors.Is(err, domain.ErrInvalidServiceName):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	resp := mapper.DomainToResp(*sub)

	data, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sub, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrSubscriptionNotFound):
			http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	resp := mapper.DomainToResp(*sub)

	data, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("user_id")
	serviceName := q.Get("service_name")
	from := q.Get("from")
	to := q.Get("to")

	var req dto.CreateSubscriptionFilterReq
	req.UserID = &userID
	req.ServiceName = &serviceName
	req.From = &from
	req.To = &to

	subs, err := h.svc.List(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data []byte
	for i := range subs {
		resp := mapper.DomainToResp(subs[i])
		data, err = json.MarshalIndent(resp, "", " ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetTotalCost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {

}
