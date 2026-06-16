-- +goose Up

-- 1. Create users table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,

    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,

    name TEXT NOT NULL,

    role TEXT NOT NULL DEFAULT 'user',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    last_login TIMESTAMP NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- 2. Add user_id to appliances (nullable first to avoid breaking existing rows)
ALTER TABLE appliances
ADD COLUMN user_id BIGINT;

-- 3. Add user_id to settings
ALTER TABLE settings
ADD COLUMN user_id BIGINT;

-- 4. Create a default admin user (for existing system)
INSERT INTO users (email, password_hash, name, role)
VALUES (
    'admin@local',
    '$2a$10$REPLACE_WITH_BCRYPT_HASH',
    'Default Admin',
    'admin'
);

-- 5. Assign existing data to this default user
-- (assuming first user = ID 1)
UPDATE appliances SET user_id = 1 WHERE user_id IS NULL;
UPDATE settings SET user_id = 1 WHERE user_id IS NULL;

-- 6. Make user_id NOT NULL after backfill
ALTER TABLE appliances
ALTER COLUMN user_id SET NOT NULL;

ALTER TABLE settings
ALTER COLUMN user_id SET NOT NULL;

-- 7. Add foreign key constraints
ALTER TABLE appliances
ADD CONSTRAINT fk_appliances_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE settings
ADD CONSTRAINT fk_settings_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

CREATE UNIQUE INDEX idx_settings_user_id ON settings(user_id);


-- +goose Down

-- Remove constraints first
ALTER TABLE appliances DROP CONSTRAINT IF EXISTS fk_appliances_user;
ALTER TABLE settings DROP CONSTRAINT IF EXISTS fk_settings_user;

-- Remove indexes
DROP INDEX IF EXISTS idx_settings_user_id;

-- Drop columns
ALTER TABLE appliances DROP COLUMN IF EXISTS user_id;
ALTER TABLE settings DROP COLUMN IF EXISTS user_id;

-- Drop users table
DROP TABLE IF EXISTS users;