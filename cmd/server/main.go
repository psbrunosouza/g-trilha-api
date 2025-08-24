package main

import (
	"log"
	"os"
	database "trilha-api/internal/shared/config"
	"trilha-api/internal/shared/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	database.ConnectDatabase()

	r := router.Router()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
