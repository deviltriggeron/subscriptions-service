package router

import (
	"subscriptions-service/internal/handler"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1.HandleFunc("/subscriptions", handler.Create).Methods("POST")
	v1.HandleFunc("/subscriptions/{id}", handler.Get).Methods("GET")
	v1.HandleFunc("/subscriptions", handler.GetAll).Methods("GET")
	v1.HandleFunc("subscriptions/total", handler.GetTotalCost).Methods("GET")
	v1.HandleFunc("/subscriptions/{id}", handler.Delete).Methods("DELETE")
	v1.HandleFunc("/subscriptions/{id}", handler.Update).Methods("PATCH")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return r
}
