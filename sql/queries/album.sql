-- CREATE TABLE "album" (
--   "id" bigserial PRIMARY KEY,
--   "name" varchar NOT NULL,
--   "image_url" varchar NOT NULL,
--   "is_deleted" bool NOT NULL DEFAULT false,
--   "singer_id" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

-- name: CreateAlbum :one
INSERT INTO album (
    name,
    image_url,
    singer_id
) VALUES (
    $1, $2, $3
) RETURNING id, name, image_url, singer_id;

-- SELECT u.*, r.name FROM users as u JOIN role as r ON u.role_id = r.id WHERE username = $1;

-- name: GetAlbum :one
SELECT a.id, a.name, a.image_url, a.singer_id, s.fullname 
FROM album as a 
JOIN singers as s ON a.singer_id = s.id  
WHERE a.id = $1 AND a.is_deleted != TRUE
LIMIT 1;

-- name: GetListAlbum :many
SELECT id FROM album
WHERE is_deleted != TRUE
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetSingerAlbums :many
SELECT id FROM album
WHERE singer_id = $1 AND is_deleted != TRUE
ORDER BY id DESC
LIMIT $2
OFFSET $3;

-- name: CountAlbum :one
SELECT count(id) FROM album;

-- name: CountSingerAlbum :one
SELECT count(id) FROM album
WHERE singer_id = $1;

-- name: UpdateAlbum :one
UPDATE album 
SET 
    name = COALESCE($2, name), 
    image_url = COALESCE($3, image_url)
WHERE id = $1
RETURNING id, name, image_url, singer_id;

-- name: DeleteAlbum :exec
UPDATE album
SET 
    is_deleted = TRUE
WHERE id = $1;
