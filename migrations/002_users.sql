-- +goose Up
CREATE TABLE admin.users (
    id bigserial PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    updated_at TIMESTAMP,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

-- +goose Down
DROP TABLE admin.users;