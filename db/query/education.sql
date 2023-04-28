-- name: GetEducations :many
SELECT *
FROM education
LIMIT 1;

-- name: CreateEducation :one
INSERT INTO
    education (name, link, description, date)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: DeleteEducation :exec
DELETE FROM education
WHERE id = $1;