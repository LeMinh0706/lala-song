package favorite

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/song"
	"github.com/google/uuid"
)

type FavoriteService struct {
	q *db.Queries
	s song.ISongService
}

// CreateLikeSong implements IFavoriteService.
func (f *FavoriteService) CreateLikeSong(ctx context.Context, username string, song_id uuid.UUID) (*db.Favorite, error) {
	user, err := f.q.GetUserId(ctx, username)
	if err != nil {
		return &db.Favorite{}, err
	}
	like, err := f.q.CreateLikeSong(ctx, db.CreateLikeSongParams{
		UserID: user,
		SongID: song_id,
	})
	if err != nil {
		return &db.Favorite{}, err
	}
	return &like, nil
}

// GetLikeSong implements IFavoriteService.
func (f *FavoriteService) GetLikeSong(ctx context.Context, username string, song_id uuid.UUID) (bool, error) {
	user, err := f.q.GetUserId(ctx, username)
	if err != nil {
		return false, nil
	}
	like, err := f.q.GetFavorite(ctx, db.GetFavoriteParams{UserID: user, SongID: song_id})
	if err != nil {
		return false, nil
	}
	if like == uuid.Nil {
		return false, nil
	}
	return true, nil
}

// GetListSong implements IFavoriteService.
func (f *FavoriteService) GetListSong(ctx context.Context, username string, page int32, pageSize int32) ([]song.SongResponse, error) {
	var res []song.SongResponse
	user, err := f.q.GetUserId(ctx, username)
	if err != nil {
		return res, err
	}
	list, err := f.q.GetFavoriteSongs(ctx, db.GetFavoriteSongsParams{
		UserID: user,
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil {
		return res, err
	}
	if len(list) == 0 {
		return []song.SongResponse{}, nil
	}
	for _, element := range list {
		song, _ := f.s.GetSong(ctx, element)
		res = append(res, song)
	}
	return res, nil
}

// Unlike implements IFavoriteService.
func (f *FavoriteService) Unlike(ctx context.Context, username string, song_id uuid.UUID) error {
	user, err := f.q.GetUserId(ctx, username)
	if err != nil {
		return err
	}

	err = f.q.UnlikeSong(ctx, db.UnlikeSongParams{UserID: user, SongID: song_id})
	if err != nil {
		return err
	}

	return nil
}

func NewFavoriteService(q *db.Queries, s song.ISongService) IFavoriteService {
	return &FavoriteService{
		q: q,
		s: s,
	}
}
