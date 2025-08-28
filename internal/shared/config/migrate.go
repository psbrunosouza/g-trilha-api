package config

import (
	"log"

	"ariga.io/atlas/sql/migrate"
)

func runMigrations() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://admin:admin@postgres:5432/trilha_app_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
