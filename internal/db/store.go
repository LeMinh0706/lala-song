package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *Store) CreateSongTx(ctx context.Context, uuid uuid.UUID, name string, song_file string, lyric_file string, album_id int64) (CreateSongRow, error) {
	var result CreateSongRow
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		singer, err := q.GetSingerAlbum(ctx, album_id)
		if err != nil {
			return err
		}

		result, err = q.CreateSong(ctx, CreateSongParams{
			ID:        uuid,
			Name:      name,
			SongFile:  song_file,
			LyricFile: lyric_file,
			AlbumID:   album_id,
			Lyrics:    "Set cứng",
		})
		if err != nil {
			return err
		}

		_, err = q.AddSongSinger(ctx, AddSongSingerParams{SingerID: singer, SongID: uuid})
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}
