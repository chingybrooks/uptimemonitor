-- +goose Up
UPDATE monitors
SET group_name = 'Ungrouped'
WHERE group_name IS NULL OR group_name = '';

-- +goose Down
UPDATE monitors
SET group_name = ''
WHERE group_name = 'Ungrouped';
