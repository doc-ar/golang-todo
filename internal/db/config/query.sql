-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY value;

-- name: CreateTodo :one
INSERT INTO todos (value)
VALUES ($1)
RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos
SET value = $2, completed = $3
WHERE id = $1;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;
