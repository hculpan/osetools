-- name: InsertUser :exec
INSERT INTO users (username, realname, password, create_datetime)
VALUES (?, ?, ?, now());

-- name: GetLatestUserByID :one
SELECT * FROM users WHERE id = LAST_INSERT_ID();

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: GetUserId :one
SELECT id
FROM users
WHERE username = ? 
LIMIT 1;

