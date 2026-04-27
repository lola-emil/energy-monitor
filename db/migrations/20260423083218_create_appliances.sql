-- +goose Up
SELECT 'up SQL query';

CREATE TABLE appliances (
    id           BIGSERIAL PRIMARY KEY,
    name         TEXT        NOT NULL,
    location     TEXT        NOT NULL DEFAULT '',
    status       TEXT        NOT NULL DEFAULT 'offline', -- 'online' | 'offline'
    last_reading TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_appliances_status ON appliances (status);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE IF EXISTS appliances;
