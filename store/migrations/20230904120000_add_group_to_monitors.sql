-- +goose Up
ALTER TABLE monitors ADD COLUMN group_name TEXT;

-- +goose Down
-- SQLite не поддерживает DROP COLUMN, так что down делаем no-op
-- (если нужно будет убрать колонку, придётся пересоздавать таблицу)
