-- name: SearchSong :many
SELECT id FROM songs
WHERE songs.name ILIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3;

-- name: SearchSongsByLyrics :many
SELECT id
FROM songs
WHERE to_tsvector('vietnamese', lyrics) @@ plainto_tsquery('vietnamese', $1)
  AND is_deleted = false
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;
