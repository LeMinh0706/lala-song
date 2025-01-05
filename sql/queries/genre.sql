-- name: CreateGenre :one
INSERT INTO genres (
    name,
    image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetGenre :one
SELECT * FROM genres 
WHERE id = $1;

-- name: GetListGenre :many
SELECT * FROM genres
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: CountGenre :one
SELECT count(id) FROM genres;

-- name: UpdateGenre :one
UPDATE genres 
SET 
    name = COALESCE($2, name), 
    image_url = COALESCE($3, image_url)
WHERE id = $1
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres WHERE id = $1;

-- name: DeleteGenreSong :exec
DELETE FROM song_genre WHERE genres_id = $1;