-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE appliances
ADD COLUMN device_code TEXT;

UPDATE appliances
SET device_code = CONCAT(
  'EMS-',
  UPPER(SUBSTRING(MD5(RANDOM()::TEXT) FROM 1 FOR 8))
)
WHERE device_code IS NULL;

ALTER TABLE appliances
ALTER COLUMN device_code SET NOT NULL;

ALTER TABLE appliances
ADD CONSTRAINT appliances_device_code_unique
UNIQUE (device_code);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE appliances
DROP CONSTRAINT IF EXISTS appliances_device_code_unique;

ALTER TABLE appliances
DROP COLUMN IF EXISTS device_code;
-- +goose StatementEnd
