-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, name, amount, list_id)
VALUES (
gen_random_uuid(),
now(),
now(),
$1,
$2,
$3
)
RETURNING *;

-- name: GetItemsForList :many
SELECT * FROM items
WHERE list_id = $1
AND removed IS NULL;

-- name: GetItemByID :one
SELECT * FROM items
WHERE id = $1
AND removed IS NULL;

-- name: GetAllItems :many
SELECT * FROM items
WHERE removed IS NULL;

-- name: DeleteItem :exec
UPDATE items
SET removed = now()
WHERE id = $1;