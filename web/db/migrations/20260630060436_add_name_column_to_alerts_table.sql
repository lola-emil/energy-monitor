-- +goose Up
SELECT 'up SQL query';
ALTER TABLE alerts ADD COLUMN name TEXT;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE alerts DROP COLUMN name;