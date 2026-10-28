-- name: CreateNote :one
INSERT INTO note (id, title, body)
VALUES (
    sqlc.arg(id),
    sqlc.arg(title),
    sqlc.arg(body)
)
RETURNING id, title, body;


-- name: GetNote :one
SELECT id, title, body
FROM note
WHERE id = sqlc.arg(id);

-- name: DeleteNote :exec
DELETE FROM note WHERE id = sqlc.arg(id);

