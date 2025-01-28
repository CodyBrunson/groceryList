-- +goose Up
CREATE TABLE items (
id UUID PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
removed TIMESTAMP,
name TEXT NOT NULL,
amount TEXT NOT NULL
);

-- +goose Down
DROP TABLE items;