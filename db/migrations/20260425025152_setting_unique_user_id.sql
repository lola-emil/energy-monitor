-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


ALTER TABLE settings
ADD CONSTRAINT settings_user_id_unique
UNIQUE (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';


ALTER TABLE settings
DROP CONSTRAINT IF EXISTS settings_user_id_unique;
-- +goose StatementEnd
