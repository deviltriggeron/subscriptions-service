package db

import (
	"database/sql"
	"fmt"
	"log"

	"subscriptions-service/internal/domain"
)

func InitDB(cfg domain.DBConfig) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pass, cfg.DB)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error init db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error ping db: %v", err)
	}

	return db
}
