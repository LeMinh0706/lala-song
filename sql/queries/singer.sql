-- name: CreateSinger :one
INSERT INTO singers (
    fullname,
    image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSinger :one
SELECT * FROM singers 
WHERE id = $1;

-- name: GetListSinger :many
SELECT * FROM singers
LIMIT $1
OFFSET $2;
