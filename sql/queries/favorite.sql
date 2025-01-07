-- name: CreateLikeSong :one
INSERT INTO favorite (
    user_id,
    song_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: UnlikeSong :exec
DELETE FROM favorite WHERE user_id = $1 AND song_id = $2;

-- name: GetFavoriteSongs :many
SELECT song_id FROM favorite WHERE user_id = $1
LIMIT $2
OFFSET $3;

-- name: GetFavorite :one
SELECT song_id FROM favorite WHERE user_id = $1 AND song_id = $2;
