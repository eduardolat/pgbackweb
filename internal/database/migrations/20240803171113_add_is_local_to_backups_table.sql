-- +goose Up
-- +goose StatementBegin
ALTER TABLE backups ALTER COLUMN destination_id DROP NOT NULL;
ALTER TABLE backups ADD COLUMN is_local BOOLEAN NOT NULL DEFAULT FALSE;
ALTER TABLE backups ADD CONSTRAINT backups_destination_check CHECK (
  (is_local = TRUE AND destination_id IS NULL) OR
  (is_local = FALSE AND destination_id IS NOT NULL)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE backups DROP CONSTRAINT backups_destination_check;
ALTER TABLE backups DROP COLUMN is_local;
ALTER TABLE backups ALTER COLUMN destination_id SET NOT NULL;
-- +goose StatementEnd
