package router

import (
	"subscriptions-service/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	v1 := mux.NewRouter()

	return v1
}
