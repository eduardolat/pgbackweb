-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS hstore;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP EXTENSION IF EXISTS hstore;
-- +goose StatementEnd
