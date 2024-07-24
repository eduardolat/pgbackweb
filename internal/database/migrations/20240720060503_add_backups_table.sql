-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS backups (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  database_id UUID NOT NULL REFERENCES databases(id) ON DELETE CASCADE,
  destination_id UUID NOT NULL REFERENCES destinations(id) ON DELETE CASCADE,

  name TEXT NOT NULL,
  cron_expression TEXT NOT NULL,
  time_zone TEXT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,
  dest_dir TEXT NOT NULL,
  retention_days SMALLINT NOT NULL DEFAULT 0,

  opt_data_only BOOLEAN NOT NULL DEFAULT FALSE,
  opt_schema_only BOOLEAN NOT NULL DEFAULT FALSE,
  opt_clean BOOLEAN NOT NULL DEFAULT FALSE,
  opt_if_exists BOOLEAN NOT NULL DEFAULT FALSE,
  opt_create BOOLEAN NOT NULL DEFAULT FALSE,
  opt_no_comments BOOLEAN NOT NULL DEFAULT FALSE,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

CREATE TRIGGER backups_change_updated_at
BEFORE UPDATE ON backups FOR EACH ROW EXECUTE FUNCTION change_updated_at();

CREATE INDEX IF NOT EXISTS
idx_backups_database_id ON backups(database_id);

CREATE INDEX IF NOT EXISTS
idx_backups_destination_id ON backups(destination_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS backups;
-- +goose StatementEnd
