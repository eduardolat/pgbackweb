-- +goose Up
-- +goose StatementBegin
ALTER TABLE databases
  DROP CONSTRAINT IF EXISTS databases_pg_version_check,
  ADD CONSTRAINT databases_pg_version_check
  CHECK (pg_version IN ('13', '14', '15', '16', '17'));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE databases
  DROP CONSTRAINT IF EXISTS databases_pg_version_check,
  ADD CONSTRAINT databases_pg_version_check
  CHECK (pg_version IN ('13', '14', '15', '16'));
-- +goose StatementEnd
