// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: album.sql

package db

import (
	"context"
)

const countAlbum = `-- name: CountAlbum :one
SELECT count(id) FROM album
`

func (q *Queries) CountAlbum(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countAlbum)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countSingerAlbum = `-- name: CountSingerAlbum :one
SELECT count(id) FROM album
WHERE singer_id = $1
`

func (q *Queries) CountSingerAlbum(ctx context.Context, singerID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, countSingerAlbum, singerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAlbum = `-- name: CreateAlbum :one

INSERT INTO album (
    name,
    image_url,
    singer_id
) VALUES (
    $1, $2, $3
) RETURNING id, name, image_url, singer_id
`

type CreateAlbumParams struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	SingerID int64  `json:"singer_id"`
}

type CreateAlbumRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	SingerID int64  `json:"singer_id"`
}

// CREATE TABLE "album" (
//
//	"id" bigserial PRIMARY KEY,
//	"name" varchar NOT NULL,
//	"image_url" varchar NOT NULL,
//	"is_deleted" bool NOT NULL DEFAULT false,
//	"singer_id" bigint NOT NULL,
//	"created_at" timestamptz NOT NULL DEFAULT (now())
//
// );
func (q *Queries) CreateAlbum(ctx context.Context, arg CreateAlbumParams) (CreateAlbumRow, error) {
	row := q.db.QueryRowContext(ctx, createAlbum, arg.Name, arg.ImageUrl, arg.SingerID)
	var i CreateAlbumRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.SingerID,
	)
	return i, err
}

const deleteAlbum = `-- name: DeleteAlbum :exec
UPDATE album
SET 
    is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) DeleteAlbum(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAlbum, id)
	return err
}

const getAlbum = `-- name: GetAlbum :one

SELECT a.id, a.name, a.image_url, a.singer_id, s.fullname 
FROM album as a 
JOIN singers as s ON a.singer_id = s.id  
WHERE a.id = $1 AND a.is_deleted != TRUE
LIMIT 1
`

type GetAlbumRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	SingerID int64  `json:"singer_id"`
	Fullname string `json:"fullname"`
}

// SELECT u.*, r.name FROM users as u JOIN role as r ON u.role_id = r.id WHERE username = $1;
func (q *Queries) GetAlbum(ctx context.Context, id int64) (GetAlbumRow, error) {
	row := q.db.QueryRowContext(ctx, getAlbum, id)
	var i GetAlbumRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.SingerID,
		&i.Fullname,
	)
	return i, err
}

const getListAlbum = `-- name: GetListAlbum :many
SELECT id FROM album
WHERE is_deleted != TRUE
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type GetListAlbumParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListAlbum(ctx context.Context, arg GetListAlbumParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getListAlbum, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
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

const getSingerAlbums = `-- name: GetSingerAlbums :many
SELECT id FROM album
WHERE singer_id = $1 AND is_deleted != TRUE
ORDER BY id DESC
LIMIT $2
OFFSET $3
`

type GetSingerAlbumsParams struct {
	SingerID int64 `json:"singer_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) GetSingerAlbums(ctx context.Context, arg GetSingerAlbumsParams) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getSingerAlbums, arg.SingerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
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

const updateAlbum = `-- name: UpdateAlbum :one
UPDATE album 
SET 
    name = COALESCE($2, name), 
    image_url = COALESCE($3, image_url)
WHERE id = $1
RETURNING id, name, image_url, singer_id
`

type UpdateAlbumParams struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type UpdateAlbumRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	SingerID int64  `json:"singer_id"`
}

func (q *Queries) UpdateAlbum(ctx context.Context, arg UpdateAlbumParams) (UpdateAlbumRow, error) {
	row := q.db.QueryRowContext(ctx, updateAlbum, arg.ID, arg.Name, arg.ImageUrl)
	var i UpdateAlbumRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.SingerID,
	)
	return i, err
}
