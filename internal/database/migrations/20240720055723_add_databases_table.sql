-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS databases (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,

  name TEXT NOT NULL UNIQUE,
  connection_string TEXT NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

CREATE TRIGGER databases_change_updated_at
BEFORE UPDATE ON databases FOR EACH ROW EXECUTE FUNCTION change_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS databases;
-- +goose StatementEnd
