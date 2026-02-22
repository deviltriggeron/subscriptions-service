package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"subscriptions-service/internal/config"
	"subscriptions-service/internal/db"
	"subscriptions-service/internal/handler"
	"subscriptions-service/internal/repository"
	"subscriptions-service/internal/router"
	"subscriptions-service/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	srv := buildServer()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Listen and runnnig :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server will shutdown gracefully...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutdown server: %v", err)
	}
	wg.Wait()
}

func buildServer() http.Server {
	cfg := config.GetDBConfig()
	db := db.InitDB(cfg)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	router := router.NewRouter(handler)

	return http.Server{
		Handler: router,
		Addr:    config.GetServerPort(),
	}
}
