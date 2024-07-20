-- name: AuthServiceDeleteOldSessions :exec
DELETE FROM sessions WHERE created_at <= @date_threshold;
