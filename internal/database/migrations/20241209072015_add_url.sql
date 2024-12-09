-- +goose Up
-- +goose StatementBegin
CREATE TABLE URL (
     id VARCHAR(255) PRIMARY KEY,
     original_url TEXT NOT NULL,
     short_url VARCHAR(255) UNIQUE NOT NULL,
     device_id VARCHAR(255),
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     expires_at TIMESTAMP,
     click_count INT NOT NULL DEFAULT 0,
     is_active BOOLEAN NOT NULL DEFAULT TRUE,
     FOREIGN KEY (device_id) REFERENCES DEVICE(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE URL;
-- +goose StatementEnd
