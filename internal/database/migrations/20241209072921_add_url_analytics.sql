-- +goose Up
-- +goose StatementBegin
CREATE TABLE URL_ANALYTICS (
   id VARCHAR(255) PRIMARY KEY,
   url_id VARCHAR(255) NOT NULL,
   device_id VARCHAR(255),
   click_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   ip_address VARCHAR(255),
   user_agent TEXT,
   geolocation VARCHAR(255),
   FOREIGN KEY (url_id) REFERENCES URL(id),
   FOREIGN KEY (device_id) REFERENCES DEVICE(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE URL_ANALYTICS;
-- +goose StatementEnd
