-- name: CreateAccount :one
INSERT INTO accounts (name, email, password, avatar)
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, password, avatar, created_at, updated_at, deleted_at;
