package song

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/google/uuid"
)

type SongService struct {
	q *db.Queries
}

// CreateSong implements ISongService.
func (s *SongService) CreateSong(ctx context.Context, uuid uuid.UUID, name string, song_file string, lyric_file string, album_id int64) (*db.CreateSongRow, error) {
	song, err := s.q.CreateSong(ctx, db.CreateSongParams{
		ID:        uuid,
		Name:      name,
		SongFile:  song_file,
		LyricFile: lyric_file,
		AlbumID:   album_id,
	})
	if err != nil {
		return &db.CreateSongRow{}, err
	}
	return &song, nil
}

// DeleteSong implements ISongService.
func (s *SongService) DeleteSong(ctx context.Context, uuid uuid.UUID) error {
	panic("unimplemented")
}

// GetListSong implements ISongService.
func (s *SongService) GetListSong(ctx context.Context, singer_id int64, album_id int64, genres int64, filter string) ([]db.GetSongRow, int, error) {
	panic("unimplemented")
}

// GetSong implements ISongService.
func (s *SongService) GetSong(ctx context.Context, uuid uuid.UUID) (SongResponse, error) {
	panic("unimplemented")
}

func NewSongService(q *db.Queries) ISongService {
	return &SongService{
		q: q,
	}
}
