package song

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/google/uuid"
)

type SongService struct {
	q *db.Queries
}

// AddGenreSong implements ISongService.
func (s *SongService) AddGenreSong(ctx context.Context, uuid uuid.UUID, genre_id int64) (*db.SongGenre, error) {
	create, err := s.q.AddSongGenre(ctx, db.AddSongGenreParams{
		GenresID: genre_id,
		SongID:   uuid,
	})
	if err != nil {
		return &db.SongGenre{}, err
	}
	return &create, nil
}

// AddFeatureSong implements ISongService.
func (s *SongService) AddFeatureSong(ctx context.Context, uuid uuid.UUID, singer_id int64) (*db.SingerSong, error) {
	create, err := s.q.AddSongSinger(ctx, db.AddSongSingerParams{
		SingerID: singer_id,
		SongID:   uuid,
	})
	if err != nil {
		return &db.SingerSong{}, err
	}
	return &create, nil
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

	singer, err := s.q.GetSingerAlbum(ctx, album_id)
	if err != nil {
		return &db.CreateSongRow{}, err
	}

	_, err = s.q.AddSongSinger(ctx, db.AddSongSingerParams{SingerID: singer, SongID: uuid})
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
	song, err := s.q.GetSong(ctx, uuid)
	if err != nil {
		return SongResponse{}, err
	}
	genres, _ := s.q.GetGenresWithSong(ctx, uuid)
	singer, _ := s.q.GetSingersWithSong(ctx, uuid)

	return SongResponse{Genres: genres, Song: song, Singer: singer}, nil
}

func NewSongService(q *db.Queries) ISongService {
	return &SongService{
		q: q,
	}
}
