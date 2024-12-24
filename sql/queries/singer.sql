-- name: CreateSinger :one
INSERT INTO singers (
    fullname,
    image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSinger :one
SELECT * FROM singers 
LIMIT $1 
OFFSET $2;