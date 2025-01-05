package song

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/google/uuid"
)

type ISongService interface {
	CreateSong(ctx context.Context, uuid uuid.UUID, name, song_file, lyric_file string, album_id int64) (*db.CreateSongRow, error)
	GetSong(ctx context.Context, uuid uuid.UUID) (SongResponse, error)
	GetListSong(ctx context.Context, singer_id, album_id, genres int64, filter string) ([]db.GetSongRow, int, error)
	DeleteSong(ctx context.Context, uuid uuid.UUID) error
}
