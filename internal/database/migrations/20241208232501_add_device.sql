-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE DEVICE (
    id VARCHAR(255) PRIMARY KEY,
    device_identifier VARCHAR(255) UNIQUE NOT NULL,
    usage_count INT NOT NULL DEFAULT 0,
    first_usage_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE DEVICE;
-- +goose StatementEnd
