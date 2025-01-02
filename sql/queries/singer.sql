-- name: CreateSinger :one
INSERT INTO singers (
    fullname,
    image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSinger :one
SELECT id, fullname, image_url FROM singers 
WHERE id = $1;

-- name: GetListSinger :many
SELECT id, fullname, image_url FROM singers
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: CountSinger :one
SELECT count(id) FROM singers;
