-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
INSERT INTO users (
    username,
    password_hash,
    name,
    role,
    is_active,
    created_at,
    updated_at
)
VALUES (
    'admin',
    '$argon2id$v=19$m=65536,t=3,p=2$SPqc3Snj6xfvpJzljSi8xg$77cfCl27aWQ438W5xc3lcXkIZLKw2treSs49szchRc4',
    'System Administrator',
    'admin',
    true,
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM users
WHERE username = 'admin';
-- +goose StatementEnd
