-- name: GetProjects :many
SELECT *
FROM projects
LIMIT 1;

-- name: CreateProject :one
INSERT INTO
    projects (name, link, description, date)
VALUES
    ($1, $2, $3, $4) RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM projects
WHERE id = $1;