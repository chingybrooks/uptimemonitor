-- +goose Up
ALTER TABLE monitors
    ALTER COLUMN group_name SET DEFAULT 'Ungrouped';

UPDATE monitors
    SET group_name = 'Ungrouped'
    WHERE group_name = '';

-- +goose Down
ALTER TABLE monitors
    ALTER COLUMN group_name SET DEFAULT '';
