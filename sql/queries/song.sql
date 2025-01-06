-- CREATE TABLE "songs" (
--   "id" uuid PRIMARY KEY,
--   "name" varchar NOT NULL,
--   "song_file" varchar NOT NULL,
--   "lyric_file" varchar NOT NULL,
--   "is_deleted" bool NOT NULL DEFAULT false,
--   "album_id" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- name: CreateSong :one
INSERT INTO songs (
    id,
    name,
    song_file,
    lyric_file,
    album_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, song_file, lyric_file, album_id, created_at;


-- name: GetSong :one
SELECT s.id, s.name, s.song_file, s.lyric_file, s.album_id, a.name, a.image_url 
FROM songs as s
JOIN album as a ON s.album_id = a.id
WHERE s.id = $1 AND s.is_deleted IS NOT TRUE;

-- name: GetSingersWithSong :many
SELECT s.id, s.fullname, s.image_url FROM singers as s 
JOIN singer_song as i ON s.id = i.singer_id
WHERE i.song_id = $1;

-- name: GetGenresWithSong :many
SELECT g.* FROM genres as g 
JOIN song_genre as s ON g.id = s.genres_id
WHERE s.song_id = $1;


-- name: GetListSong :many
SELECT id FROM songs 
WHERE is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $1
OFFSET $2;

-- name: GetAlbumSongs :many
SELECT s.id FROM songs as s 
JOIN album as a ON s.album_id = a.id 
WHERE album_id = $1 AND is_deleted != TRUE
ORDER BY s.created_at DESC 
LIMIT $2
OFFSET $3;

-- name: GetSingerSongs :many
SELECT id FROM songs as s 
JOIN singer_song as i ON s.id = i.song_id 
WHERE i.singer_id = $1 AND is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $2
OFFSET $3;

-- name: GetGenreSongs :many
SELECT id FROM songs as s 
JOIN song_genre as g ON s.id = g.song_id 
WHERE g.genres_id = $1 AND is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $1
OFFSET $2;

-- name: DeleteSong :exec
UPDATE songs
SET 
    is_deleted = TRUE
WHERE id = $1;


-- name: AddSongGenre :one
INSERT INTO song_genre (
    genres_id,
    song_id
)VALUES(
    $1, $2
)RETURNING *;


-- name: AddSongSinger :one
INSERT INTO singer_song (
    singer_id,
    song_id
)VALUES(
    $1, $2
)RETURNING *;

-- name: GetSingerAlbum :one
SELECT s.id FROM singers as s 
JOIN album as a ON s.id = a.singer_id
WHERE a.id = $1; 