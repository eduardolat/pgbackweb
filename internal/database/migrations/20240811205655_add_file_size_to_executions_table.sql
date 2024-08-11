-- +goose Up
-- +goose StatementBegin
ALTER TABLE executions ADD COLUMN file_size BIGINT NULL DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE executions DROP COLUMN file_size;
-- +goose StatementEnd
