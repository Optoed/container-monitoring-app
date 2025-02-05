CREATE TABLE IF NOT EXISTS containers (
    id SERIAL PRIMARY KEY,
    ip VARCHAR(31) NOT NULL,
    status VARCHAR(31) NOT NULL,
    last_ping_time TIMESTAMP NOT NULL,
    ping_duration INTERVAL NOT NULL
);

CREATE TABLE IF NOT EXISTS container_ping_history (
    id SERIAL PRIMARY KEY,
    container_id INTEGER REFERENCES containers(id) ON DELETE CASCADE,
    ip VARCHAR(31) NOT NULL,
    status VARCHAR(31) NOT NULL,
    ping_time TIMESTAMP NOT NULL,
    ping_duration INTERVAL NOT NULL
)