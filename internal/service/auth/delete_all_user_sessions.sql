-- name: AuthServiceDeleteAllUserSessions :exec
DELETE FROM sessions WHERE user_id = @user_id;
