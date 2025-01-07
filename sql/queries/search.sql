-- name: SearchSong :many
SELECT id FROM songs
WHERE fullname ILIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3;

-- name: SearchSongsByLyrics :many
SELECT id, name, song_file, lyric_file, lyrics, is_deleted, album_id, created_at
FROM songs
WHERE lyrics_tsv @@ plainto_tsquery('vietnamese', $1)
  AND is_deleted = false
ORDER BY created_at DESC;
