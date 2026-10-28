
-- +goose Up
CREATE TABLE IF NOT EXISTS note (
    id TEXT PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    body VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE note;