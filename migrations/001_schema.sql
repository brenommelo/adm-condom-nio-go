-- +goose Up
CREATE SCHEMA admin AUTHORIZATION postgres;

-- +goose Down
DROP SCHEMA admin;