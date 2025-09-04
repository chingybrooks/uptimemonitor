-- +goose Up
ALTER TABLE monitors ADD COLUMN group_name TEXT DEFAULT '';

-- +goose Down
ALTER TABLE monitors DROP COLUMN group_name;
