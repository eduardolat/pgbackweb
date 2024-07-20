-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION change_updated_at() RETURNS TRIGGER AS $$
  BEGIN
    IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
      NEW.updated_at = now(); 
      RETURN NEW;
    ELSE
      RETURN OLD;
    END IF;
  END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS change_updated_at();
-- +goose StatementEnd
