-- name: Register :one
INSERT INTO users(
    id,
    username, 
    password,
    fullname,
    gender,
    avt,
    role_id
) VALUES (
    $1, $2, $3, $4, $5, $6, 2
) RETURNING *;

-- name: Login :one
SELECT u.*, r.name FROM users as u JOIN role as r ON u.role_id = r.id WHERE username = $1;

-- name: GetMe :one
SELECT fullname, gender, avt, role_id FROM users
WHERE username = $1;