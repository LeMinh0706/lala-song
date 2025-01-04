-- name: CreateSinger :one
INSERT INTO singers (
    fullname,
    image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSinger :one
SELECT id, fullname, image_url FROM singers 
WHERE id = $1 AND is_deleted != TRUE;

-- name: GetListSinger :many
SELECT id, fullname, image_url FROM singers
WHERE is_deleted != TRUE
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: CountSinger :one
SELECT count(id) FROM singers WHERE is_deleted != TRUE;

-- name: UpdateSinger :one
UPDATE singers 
SET 
    fullname = COALESCE($2, fullname), 
    image_url = COALESCE($3, image_url)
WHERE id = $1
RETURNING id, fullname, image_url;

-- name: DeleteSinger :exec
UPDATE singers SET is_deleted = TRUE WHERE id = $1;