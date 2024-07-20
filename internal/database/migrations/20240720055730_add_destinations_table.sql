-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS destinations (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,

  name TEXT NOT NULL UNIQUE,
  bucket_name TEXT NOT NULL,
  access_key TEXT NOT NULL,
  secret_key TEXT NOT NULL,
  region TEXT NOT NULL,
  endpoint TEXT NOT NULL,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

CREATE TRIGGER destinations_change_updated_at
BEFORE UPDATE ON destinations FOR EACH ROW EXECUTE FUNCTION change_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS destinations;
-- +goose StatementEnd
