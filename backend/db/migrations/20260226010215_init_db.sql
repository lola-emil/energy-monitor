-- +goose Up
SELECT 'up SQL query';

-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'user_roles'
    ) THEN
        CREATE TYPE user_roles AS ENUM ('admin', 'user');
    END IF;
END
$$;
-- +goose StatementEnd

CREATE TABLE users (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,

    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,

    user_role user_roles DEFAULT 'user',

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE devices (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    device_name VARCHAR(50),
    device_serial VARCHAR(16) NOT NULL,

    is_active BOOLEAN DEFAULT FALSE,
    last_active TIMESTAMPTZ,

    activation_code VARCHAR(50) NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE device_claims (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    device_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (device_id) REFERENCES devices (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE energy_readings (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    
    device_id INTEGER NOT NULL,
    voltage NUMERIC(12, 2) DEFAULT 0,
    current NUMERIC(12, 2) DEFAULT 0,
    power_kwh NUMERIC(12, 2) DEFAULT 0,

    stamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (device_id) REFERENCES devices (id)
);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE users CASCADE;
DROP TABLE devices CASCADE;
DROP TABLE device_claims CASCADE;
DROP TABLE energy_readings CASCADE;


DROP TYPE user_roles CASCADE;
