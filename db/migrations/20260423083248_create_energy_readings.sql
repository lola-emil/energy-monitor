-- +goose Up
SELECT 'up SQL query';
CREATE TABLE energy_readings (
    id           BIGSERIAL PRIMARY KEY,
    appliance_id BIGINT     NOT NULL REFERENCES appliances(id) ON DELETE CASCADE,
    ts           TIMESTAMPTZ NOT NULL,          -- reading timestamp

    voltage      DOUBLE PRECISION NOT NULL,     -- V
    current      DOUBLE PRECISION NOT NULL,     -- A
    power        DOUBLE PRECISION NOT NULL,     -- W (instantaneous / avg over 2s)

    -- incremental energy for this interval only (e.g., 2 seconds)
    energy_kwh   DOUBLE PRECISION NOT NULL,     -- kWh for this slice

    frequency_hz DOUBLE PRECISION NOT NULL      -- Hz
);

CREATE INDEX idx_energy_readings_appliance_ts
    ON energy_readings (appliance_id, ts);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS energy_readings;