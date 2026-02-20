package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"subscriptions-service/internal/domain"
)

var once sync.Once

func loadEnv() {
	once.Do(func() {
		if err := godotenv.Load("config.env"); err != nil {
			log.Println("config.env not found, using system env")
		}
	})
}

func GetDBConfig() domain.DBConfig {
	loadEnv()

	return domain.DBConfig{
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		DB:   os.Getenv("POSTGRES_DB"),
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
	}
}

func GetServerPort() string {
	loadEnv()
	return os.Getenv("SERVER_PORT")
}
