-- name: CreateUser :exec
INSERT INTO users (user_id, username)
VALUES ($1, $2)
    ON CONFLICT (user_id, username)
DO UPDATE SET username = EXCLUDED.username;

-- name: ApproveCheck :exec
update users set purchased = TRUE where user_id = $1;

-- name: GetLanguage :one
SELECT COALESCE(language_code, 'en') AS language_code FROM users WHERE user_id = $1;

-- name: ChangeLanguage :exec
update users set language_code = $2 where user_id = $1;


