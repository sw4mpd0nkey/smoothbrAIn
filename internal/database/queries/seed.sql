

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: InsertSeedUser :exec
INSERT INTO users (email, display_name, role, password_hash)
VALUES (?, ?, ?, ?);

:w
