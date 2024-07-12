-- name: InsertUser :one
INSERT INTO users (username, realname, password, create_datetime)
VALUES (?, ?, ?, datetime('now', 'localtime'))
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: GetAllCharacters :many
SELECT c.*, u.username
FROM characters c, users u
WHERE c.user_id = u.user_id;

-- name: GetUserId :one
SELECT id
FROM users
WHERE username = ? 
LIMIT 1;

