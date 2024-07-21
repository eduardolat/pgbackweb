-- name: AuthServiceGetUserByToken :one
SELECT
  users.*,
  sessions.id as session_id
FROM sessions
JOIN users ON users.id = sessions.user_id
WHERE pgp_sym_decrypt(sessions.token, @encryption_key) = @token;
