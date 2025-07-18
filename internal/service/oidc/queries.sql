-- name: OIDCServiceCreateUser :one
INSERT INTO users (name, email, oidc_provider, oidc_subject)
VALUES (@name, lower(@email), @oidc_provider, @oidc_subject)
RETURNING *;

-- name: OIDCServiceGetUserByOIDC :one
SELECT * FROM users 
WHERE oidc_provider = @oidc_provider AND oidc_subject = @oidc_subject;

-- name: OIDCServiceGetUserByEmail :one
SELECT * FROM users 
WHERE email = lower(@email);

-- name: OIDCServiceUpdateUser :one
UPDATE users 
SET name = @name, email = lower(@email), updated_at = NOW()
WHERE oidc_provider = @oidc_provider AND oidc_subject = @oidc_subject
RETURNING *;
