package favorite

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/song"
	"github.com/google/uuid"
)

type IFavoriteService interface {
	CreateLikeSong(ctx context.Context, username string, song_id uuid.UUID) (*db.Favorite, error)
	Unlike(ctx context.Context, username string, song_id uuid.UUID) error
	GetListSong(ctx context.Context, username string, page, pageSize int32) ([]song.SongResponse, error)
	GetLikeSong(ctx context.Context, username string, song_id uuid.UUID) (bool, error)
}
