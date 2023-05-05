-- name: GetUser :one
SELECT *
FROM admin
WHERE login = $1
LIMIT 1;