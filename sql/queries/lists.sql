-- name: GetAllLists :many
SELECT * FROM lists
WHERE deleted_at IS NULL;

-- name: CreateList :one
INSERT INTO lists (id, created_at, updated_at, name)
VALUES (
gen_random_uuid(),
now(),
now(),
$1
)
RETURNING *;

-- name: GetListByID :one
SELECT * FROM lists
WHERE id = $1
AND deleted_at IS NULL;

-- name: DeleteListByID :exec
UPDATE lists
SET deleted_at = now()
WHERE id = $1;