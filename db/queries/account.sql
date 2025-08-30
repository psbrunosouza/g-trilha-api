-- name: CreateAccount :one
INSERT INTO accounts (name, email, password, avatar)
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, password, avatar, created_at, updated_at, deleted_at;

-- name: UpdateAccount :one
UPDATE accounts
SET name = $2, avatar = $3
WHERE id = $1
RETURNING id, name, email, password, avatar, created_at, updated_at, deleted_at;

-- name: FindAccount :one
SELECT id, name, email, avatar, created_at, updated_at, deleted_at, password
FROM accounts
WHERE id = $1;
