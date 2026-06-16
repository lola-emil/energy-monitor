-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

INSERT INTO settings (
    user_id,
    currency,
    rate_per_kwh,
    fixed_monthly_charge,
    default_analytics_range,
    refresh_interval_seconds,
    time_format,
    enable_voltage_alerts,
    over_voltage_threshold,
    under_voltage_threshold,
    enable_current_alerts,
    over_current_threshold,
    enable_offline_alerts,
    updated_at
)
SELECT
    u.id,
    'PHP',
    12.5,
    150,
    'month',
    30,
    '24h',
    TRUE,
    240,
    200,
    TRUE,
    15,
    TRUE,
    NOW()
FROM users u
WHERE u.username = 'admin'
ON CONFLICT (user_id) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM settings
WHERE user_id = (
    SELECT id FROM users WHERE username = 'admin'
);
-- +goose StatementEnd
