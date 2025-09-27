-- +goose Up
-- +goose StatementBegin
ALTER TABLE users 
ADD COLUMN oidc_provider TEXT,
ADD COLUMN oidc_subject TEXT,
ADD COLUMN password_nullable TEXT;

-- Make password nullable and copy existing passwords
UPDATE users SET password_nullable = password;
ALTER TABLE users DROP COLUMN password;
ALTER TABLE users RENAME COLUMN password_nullable TO password;

-- Create unique index for OIDC users
CREATE UNIQUE INDEX users_oidc_provider_subject_idx 
ON users (oidc_provider, oidc_subject) 
WHERE oidc_provider IS NOT NULL AND oidc_subject IS NOT NULL;

-- Add constraint to ensure either password or OIDC is set
ALTER TABLE users ADD CONSTRAINT users_auth_method_check 
CHECK (
    (password IS NOT NULL AND oidc_provider IS NULL AND oidc_subject IS NULL) OR
    (password IS NULL AND oidc_provider IS NOT NULL AND oidc_subject IS NOT NULL)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS users_oidc_provider_subject_idx;
ALTER TABLE users DROP CONSTRAINT IF EXISTS users_auth_method_check;
ALTER TABLE users DROP COLUMN IF EXISTS oidc_provider;
ALTER TABLE users DROP COLUMN IF EXISTS oidc_subject;

-- Make password required again (this will fail if there are OIDC users)
ALTER TABLE users ALTER COLUMN password SET NOT NULL;
-- +goose StatementEnd
