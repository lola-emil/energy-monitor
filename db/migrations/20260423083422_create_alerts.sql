-- +goose Up
SELECT 'up SQL query';
CREATE TABLE alerts (
    id            BIGSERIAL PRIMARY KEY,
    appliance_id  BIGINT REFERENCES appliances(id) ON DELETE CASCADE,
    type          TEXT    NOT NULL,  -- 'over_voltage' | 'under_voltage' | 'over_current' | 'offline' | ...
    severity      TEXT    NOT NULL,  -- 'info' | 'medium' | 'high'
    message       TEXT    NOT NULL,

    triggered_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    resolved_at   TIMESTAMPTZ
);

CREATE INDEX idx_alerts_appliance_ts
    ON alerts (appliance_id, triggered_at);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS alerts;
