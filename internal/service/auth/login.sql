-- name: AuthServiceLoginGetUserByEmail :one
SELECT * FROM users WHERE email = @email;

-- name: AuthServiceLoginCreateSession :one
INSERT INTO sessions (
  user_id, token, ip, user_agent
) VALUES (
  @user_id, @token, @ip, @user_agent
) RETURNING *;
