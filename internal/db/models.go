// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Album struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	IsDeleted bool      `json:"is_deleted"`
	SingerID  int64     `json:"singer_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Favorite struct {
	UserID uuid.UUID `json:"user_id"`
	SongID uuid.UUID `json:"song_id"`
}

type Genre struct {
	ID       int64          `json:"id"`
	Name     sql.NullString `json:"name"`
	ImageUrl string         `json:"image_url"`
}

type Role struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Singer struct {
	ID        int64  `json:"id"`
	Fullname  string `json:"fullname"`
	ImageUrl  string `json:"image_url"`
	IsDeleted bool   `json:"is_deleted"`
}

type SingerSong struct {
	SongID   uuid.UUID `json:"song_id"`
	SingerID int64     `json:"singer_id"`
}

type Song struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	SongFile  string    `json:"song_file"`
	LyricFile string    `json:"lyric_file"`
	IsDeleted bool      `json:"is_deleted"`
	AlbumID   int64     `json:"album_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SongGenre struct {
	GenresID int64     `json:"genres_id"`
	SongID   uuid.UUID `json:"song_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	Gender    int32     `json:"gender"`
	Avt       string    `json:"avt"`
	RoleID    int32     `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}