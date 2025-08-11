-- name: GetUser :one
SELECT * FROM users WHERE uuid = ?;

-- name: CreateUser :execresult
INSERT INTO users(name, email) VALUES (?, ?);
