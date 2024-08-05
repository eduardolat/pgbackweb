-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS restorations (
  id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
  execution_id UUID NOT NULL REFERENCES executions(id) ON DELETE CASCADE,
  database_id UUID REFERENCES databases(id) ON DELETE CASCADE,

  status TEXT NOT NULL CHECK (
    status IN ('running', 'success', 'failed')
  ) DEFAULT 'running',
  message TEXT,

  started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  finished_at TIMESTAMPTZ
);

CREATE TRIGGER restorations_change_updated_at
BEFORE UPDATE ON restorations FOR EACH ROW EXECUTE FUNCTION change_updated_at();

CREATE INDEX IF NOT EXISTS
idx_restorations_execution_id ON restorations(execution_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS restorations;
-- +goose StatementEnd
