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

    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE devices (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    device_code VARCHAR(16) NOT NULL,

    is_active BOOLEAN DEFAULT FALSE,
    last_active TIMESTAMPTZ,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE device_claims (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    device_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    device_name VARCHAR(50),


    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (device_id) REFERENCES devices (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE readings_raw (
    device_id INTEGER NOT NULL,
    bucket TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    voltage NUMERIC(12, 2) DEFAULT 0,
    current NUMERIC(12, 2) DEFAULT 0,
    power_kwh NUMERIC(12, 2) DEFAULT 0,


    PRIMARY KEY (device_id, bucket),

    FOREIGN KEY (device_id) REFERENCES devices (id)
);
 
CREATE TABLE readings_1h (
    device_id INTEGER NOT NULL,
    bucket TIMESTAMP NOT NULL,

    voltage NUMERIC(12, 2) DEFAULT 0,
    current NUMERIC(12, 2) DEFAULT 0,
    power_kwh NUMERIC(12, 2) DEFAULT 0,
    PRIMARY KEY (device_id, bucket)
);

INSERT INTO users (username, password) VALUES ('admin', '$argon2id$v=19$m=65536,t=3,p=2$SPqc3Snj6xfvpJzljSi8xg$77cfCl27aWQ438W5xc3lcXkIZLKw2treSs49szchRc4');

-- +goose Down
SELECT 'down SQL query';

DROP TABLE users CASCADE;
DROP TABLE devices CASCADE;
DROP TABLE device_claims CASCADE;
DROP TABLE energy_readings CASCADE;


DROP TYPE user_roles CASCADE;
