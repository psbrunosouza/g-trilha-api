include .env

run:
	go run ./cmd/server/main.go

build:
	CGO_ENABLED=0 go build -v -o bin/app ./cmd/server

migrate-create:
ifndef name
	$(error "Você precisa passar o nome da migration: make migrate-create name=<nome>")
endif
	@migrate create -ext=sql -dir=$(MIGRATE_PATH) -seq $(name)

migrate-up:
	@migrate -path=$(MIGRATE_PATH) -database "$(DB_URL)" -verbose up

migrate-down:
	@migrate -path=$(MIGRATE_PATH) -database "$(DB_URL)" -verbose down
