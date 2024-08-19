-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS webhooks (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,

  name TEXT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,

  event_type TEXT NOT NULL CHECK (event_type IN (
    'database_healthy', 'database_unhealthy',
    'destination_healthy', 'destination_unhealthy',
    'execution_success', 'execution_failed'
  )),
  target_ids UUID[] NOT NULL, -- database_id, restoration_id, etc.

  url TEXT NOT NULL,
  method TEXT NOT NULL CHECK (method IN ('GET', 'POST')),
  headers TEXT, -- user-defined headers in JSON format
  body TEXT, -- user-defined body in JSON format

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ
);

CREATE TRIGGER webhooks_change_updated_at
BEFORE UPDATE ON webhooks FOR EACH ROW EXECUTE FUNCTION change_updated_at();

CREATE INDEX IF NOT EXISTS
idx_webhooks_target_ids ON webhooks USING GIN (target_ids);

CREATE TABLE IF NOT EXISTS webhook_results (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  webhook_id UUID NOT NULL REFERENCES webhooks(id) ON DELETE CASCADE,

  req_method TEXT CHECK (req_method IN ('GET', 'POST')),
  req_headers TEXT,
  req_body TEXT,

  res_status SMALLINT,
  res_headers TEXT,
  res_body TEXT,
  res_duration INTEGER, -- in milliseconds

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS
idx_webhook_results_webhook_id ON webhook_results(webhook_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS webhook_results;
DROP TABLE IF EXISTS webhooks;
-- +goose StatementEnd
