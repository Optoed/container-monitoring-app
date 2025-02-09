CREATE TABLE IF NOT EXISTS containers (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(31) NOT NULL,
    status VARCHAR(31) NOT NULL,
    last_ping_time TIMESTAMP NOT NULL,
    ping_duration BIGINT
);