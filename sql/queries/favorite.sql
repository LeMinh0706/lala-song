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
SELECT id FROM songs as S JOIN favorite as f ON s.id = f.song_id WHERE user_id = $1 AND is_deleted != TRUE
LIMIT $2
OFFSET $3;

-- name: GetFavorite :one
SELECT song_id FROM favorite WHERE user_id = $1 AND song_id = $2;
