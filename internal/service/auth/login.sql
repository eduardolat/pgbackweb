-- name: AuthServiceLoginGetUserByEmail :one
SELECT * FROM users WHERE email = @email;

-- name: AuthServiceLoginCreateSession :one
INSERT INTO sessions (
  user_id, token, ip, user_agent
) VALUES (
  @user_id, pgp_sym_encrypt(@token::TEXT, @encryption_key::TEXT), @ip, @user_agent
) RETURNING *, pgp_sym_decrypt(token, @encryption_key::TEXT) AS decrypted_token;
