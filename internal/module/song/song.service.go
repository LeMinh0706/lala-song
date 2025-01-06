package song

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/google/uuid"
)

type ISongService interface {
	CreateSong(ctx context.Context, uuid uuid.UUID, name, song_file, lyric_file string, album_id int64) (*db.CreateSongRow, error)
	AddFeatureSong(ctx context.Context, uuid uuid.UUID, singer_id int64) (*db.SingerSong, error)
	AddGenreSong(ctx context.Context, uuid uuid.UUID, genre_id int64) (*db.SongGenre, error)
	GetSong(ctx context.Context, uuid uuid.UUID) (SongResponse, error)
	GetListSong(ctx context.Context, singer, album, genres, filter string, page, pageSize int32) ([]SongResponse, int, error)
	DeleteSong(ctx context.Context, uuid uuid.UUID) error
}
