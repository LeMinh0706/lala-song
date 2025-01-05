// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: song.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSong = `-- name: CreateSong :one

INSERT INTO songs (
    id,
    name,
    song_file,
    lyric_file,
    album_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, name, song_file, lyric_file, album_id, created_at
`

type CreateSongParams struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SongFile  string    `json:"song_file"`
	LyricFile string    `json:"lyric_file"`
	AlbumID   int64     `json:"album_id"`
}

type CreateSongRow struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SongFile  string    `json:"song_file"`
	LyricFile string    `json:"lyric_file"`
	AlbumID   int64     `json:"album_id"`
	CreatedAt time.Time `json:"created_at"`
}

// CREATE TABLE "songs" (
//
//	"id" uuid PRIMARY KEY,
//	"name" varchar NOT NULL,
//	"song_file" varchar NOT NULL,
//	"lyric_file" varchar NOT NULL,
//	"is_deleted" bool NOT NULL DEFAULT false,
//	"album_id" bigint NOT NULL,
//	"created_at" timestamptz NOT NULL DEFAULT (now())
//
// );
func (q *Queries) CreateSong(ctx context.Context, arg CreateSongParams) (CreateSongRow, error) {
	row := q.db.QueryRowContext(ctx, createSong,
		arg.ID,
		arg.Name,
		arg.SongFile,
		arg.LyricFile,
		arg.AlbumID,
	)
	var i CreateSongRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.SongFile,
		&i.LyricFile,
		&i.AlbumID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSong = `-- name: DeleteSong :exec
UPDATE songs
SET 
    is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) DeleteSong(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSong, id)
	return err
}

const getAlbumSongs = `-- name: GetAlbumSongs :many
SELECT s.id FROM songs as s 
JOIN album as a ON s.album_id = a.id 
WHERE album_id = $1 AND is_deleted != TRUE
ORDER BY s.created_at DESC 
LIMIT $2
OFFSET $3
`

type GetAlbumSongsParams struct {
	AlbumID int64 `json:"album_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) GetAlbumSongs(ctx context.Context, arg GetAlbumSongsParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getAlbumSongs, arg.AlbumID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGenreSongs = `-- name: GetGenreSongs :many
SELECT id FROM songs as s 
JOIN song_genre as g ON s.id = g.song_id 
WHERE g.genres_id = $1 AND is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $1
OFFSET $2
`

type GetGenreSongsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetGenreSongs(ctx context.Context, arg GetGenreSongsParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getGenreSongs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGenresWithSong = `-- name: GetGenresWithSong :many
SELECT g.id, g.name, g.image_url FROM genres as g 
JOIN song_genre as s ON g.id = s.genres_id
WHERE s.song_id = $1
LIMIT $2
OFFSET $3
`

type GetGenresWithSongParams struct {
	SongID uuid.UUID `json:"song_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) GetGenresWithSong(ctx context.Context, arg GetGenresWithSongParams) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, getGenresWithSong, arg.SongID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Genre{}
	for rows.Next() {
		var i Genre
		if err := rows.Scan(&i.ID, &i.Name, &i.ImageUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListSong = `-- name: GetListSong :many
SELECT id FROM songs 
WHERE is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $1
OFFSET $2
`

type GetListSongParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListSong(ctx context.Context, arg GetListSongParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getListSong, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSingerSongs = `-- name: GetSingerSongs :many
SELECT id FROM songs as s 
JOIN singer_song as i ON s.id = i.song_id 
WHERE i.singer_id = $1 AND is_deleted != TRUE
ORDER BY created_at DESC 
LIMIT $2
OFFSET $3
`

type GetSingerSongsParams struct {
	SingerID int64 `json:"singer_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetSingerSongs(ctx context.Context, arg GetSingerSongsParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getSingerSongs, arg.SingerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSingersWithSong = `-- name: GetSingersWithSong :many
SELECT s.id, s.fullname, s.image_url FROM singers as s 
JOIN singer_song as i ON s.id = i.singer_id
WHERE i.song_id = $1
LIMIT $2
OFFSET $3
`

type GetSingersWithSongParams struct {
	SongID uuid.UUID `json:"song_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetSingersWithSongRow struct {
	ID       int64  `json:"id"`
	Fullname string `json:"fullname"`
	ImageUrl string `json:"image_url"`
}

func (q *Queries) GetSingersWithSong(ctx context.Context, arg GetSingersWithSongParams) ([]GetSingersWithSongRow, error) {
	rows, err := q.db.QueryContext(ctx, getSingersWithSong, arg.SongID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetSingersWithSongRow{}
	for rows.Next() {
		var i GetSingersWithSongRow
		if err := rows.Scan(&i.ID, &i.Fullname, &i.ImageUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSong = `-- name: GetSong :one
SELECT s.id, s.name, s.song_file, s.lyric_file, s.album_id, a.name, a.image_url 
FROM songs as s
JOIN album as a ON s.album_id = a.id
WHERE s.id = $1
`

type GetSongRow struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SongFile  string    `json:"song_file"`
	LyricFile string    `json:"lyric_file"`
	AlbumID   int64     `json:"album_id"`
	Name_2    string    `json:"name_2"`
	ImageUrl  string    `json:"image_url"`
}

func (q *Queries) GetSong(ctx context.Context, id uuid.UUID) (GetSongRow, error) {
	row := q.db.QueryRowContext(ctx, getSong, id)
	var i GetSongRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.SongFile,
		&i.LyricFile,
		&i.AlbumID,
		&i.Name_2,
		&i.ImageUrl,
	)
	return i, err
}
