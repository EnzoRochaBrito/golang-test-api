-- name: CreateNote :one
INSERT INTO note (id, title, body)
VALUES ($1, $2, $3)
RETURNING id, title, body;


-- name: GetNote :one
SELECT id, title, body
FROM note
WHERE id = $1;

-- name: DeleteNote :exec
DELETE FROM note WHERE id = $1;