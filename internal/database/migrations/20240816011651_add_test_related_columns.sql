-- +goose Up
-- +goose StatementBegin
ALTER TABLE databases
ADD COLUMN IF NOT EXISTS test_ok BOOLEAN,
ADD COLUMN IF NOT EXISTS test_error TEXT,
ADD COLUMN IF NOT EXISTS last_test_at TIMESTAMPTZ;

ALTER TABLE destinations
ADD COLUMN IF NOT EXISTS test_ok BOOLEAN,
ADD COLUMN IF NOT EXISTS test_error TEXT,
ADD COLUMN IF NOT EXISTS last_test_at TIMESTAMPTZ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE databases
DROP COLUMN IF EXISTS test_ok,
DROP COLUMN IF EXISTS test_error,
DROP COLUMN IF EXISTS last_test_at;

ALTER TABLE destinations
DROP COLUMN IF EXISTS test_ok,
DROP COLUMN IF EXISTS test_error,
DROP COLUMN IF EXISTS last_test_at;
-- +goose StatementEnd
