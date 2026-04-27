-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- 1. Add username column (nullable first)
ALTER TABLE users
ADD COLUMN username TEXT;

-- 2. Backfill username using email
UPDATE users
SET username = LOWER(email)
WHERE username IS NULL;

-- 3. Make username NOT NULL
ALTER TABLE users
ALTER COLUMN username SET NOT NULL;

-- 4. Add unique constraint
ALTER TABLE users
ADD CONSTRAINT users_username_unique UNIQUE (username);
ALTER TABLE users DROP COLUMN email;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

-- Recreate email column if dropped (optional)
-- ALTER TABLE users ADD COLUMN email TEXT;

-- Drop constraint
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_username_unique;

-- Drop username column
ALTER TABLE users DROP COLUMN IF EXISTS username;
-- +goose StatementEnd
