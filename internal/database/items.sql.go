// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: items.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (id, created_at, updated_at, name, amount, list_id)
VALUES (
gen_random_uuid(),
now(),
now(),
$1,
$2,
$3
)
RETURNING id, created_at, updated_at, removed, name, amount, list_id
`

type CreateItemParams struct {
	Name   string
	Amount string
	ListID uuid.UUID
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem, arg.Name, arg.Amount, arg.ListID)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Removed,
		&i.Name,
		&i.Amount,
		&i.ListID,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
UPDATE items
SET removed = now()
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const getAllItems = `-- name: GetAllItems :many
SELECT id, created_at, updated_at, removed, name, amount, list_id FROM items
WHERE removed IS NULL
`

func (q *Queries) GetAllItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getAllItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Removed,
			&i.Name,
			&i.Amount,
			&i.ListID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getItemByID = `-- name: GetItemByID :one
SELECT id, created_at, updated_at, removed, name, amount, list_id FROM items
WHERE id = $1
AND removed IS NULL
`

func (q *Queries) GetItemByID(ctx context.Context, id uuid.UUID) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemByID, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Removed,
		&i.Name,
		&i.Amount,
		&i.ListID,
	)
	return i, err
}

const getItemsForList = `-- name: GetItemsForList :many
SELECT id, created_at, updated_at, removed, name, amount, list_id FROM items
WHERE list_id = $1
AND removed IS NULL
`

func (q *Queries) GetItemsForList(ctx context.Context, listID uuid.UUID) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItemsForList, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Removed,
			&i.Name,
			&i.Amount,
			&i.ListID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
