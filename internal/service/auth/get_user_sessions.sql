-- name: AuthServiceGetUserSessions :many
SELECT * FROM sessions WHERE user_id = @user_id ORDER BY created_at DESC;
