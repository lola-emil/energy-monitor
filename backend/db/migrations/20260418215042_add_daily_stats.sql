-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

-- +goose StatementEnd
CREATE TABLE daily_stats (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    avg_voltage NUMERIC(12, 2) DEFAULT 0,
    avg_current NUMERIC(12, 2) DEFAULT 0,
    peak_power NUMERIC(12, 2) DEFAULT 0,
    total_energy NUMERIC(12, 2) DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE daily_stats CASCADE;

-- +goose StatementBegin
SELECT
    'down SQL query';

-- +goose StatementEnd