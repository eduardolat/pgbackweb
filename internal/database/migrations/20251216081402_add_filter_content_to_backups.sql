-- +goose Up
-- +goose StatementBegin
ALTER TABLE backups ADD COLUMN filter_content TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE backups DROP COLUMN filter_content;
-- +goose StatementEnd
