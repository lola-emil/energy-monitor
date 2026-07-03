-- +goose Up
SELECT 'up SQL query';
CREATE TABLE settings (
    id BIGSERIAL PRIMARY KEY, -- you can keep this as 1 for single-user system

    -- Billing
    currency             TEXT            NOT NULL DEFAULT 'PHP',
    rate_per_kwh         DOUBLE PRECISION NOT NULL DEFAULT 0,
    fixed_monthly_charge DOUBLE PRECISION NOT NULL DEFAULT 0,

    -- Display
    default_analytics_range  TEXT NOT NULL DEFAULT 'month', -- 'today' | '7d' | 'month'
    refresh_interval_seconds INTEGER NOT NULL DEFAULT 5,
    time_format              TEXT NOT NULL DEFAULT '24h',   -- '24h' | '12h'

    -- Voltage alerts
    enable_voltage_alerts   BOOLEAN NOT NULL DEFAULT TRUE,
    over_voltage_threshold  DOUBLE PRECISION NOT NULL DEFAULT 250,
    under_voltage_threshold DOUBLE PRECISION NOT NULL DEFAULT 190,

    -- Current / load alerts
    enable_current_alerts   BOOLEAN NOT NULL DEFAULT TRUE,
    over_current_threshold  DOUBLE PRECISION NOT NULL DEFAULT 15,

    -- Offline alerts
    enable_offline_alerts   BOOLEAN NOT NULL DEFAULT TRUE,

    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS settings;