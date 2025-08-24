package database

import (
	"context"
	"fmt"
	"log"
	"os"
	db "trilha-api/internal/shared/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *db.Queries

func ConnectDatabase() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	DB = db.New(pool)
}
