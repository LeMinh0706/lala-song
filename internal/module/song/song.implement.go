package song

import (
	"context"
	"strconv"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/google/uuid"
)

type SongService struct {
	q  *db.Queries
	st *db.Store
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
	song, err := s.st.CreateSongTx(ctx, uuid, name, song_file, lyric_file, album_id)
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
func (s *SongService) GetListSong(ctx context.Context, singer string, album string, genres string, filter string, page int32, pageSize int32) ([]SongResponse, int, error) {
	var res []SongResponse
	switch filter {
	case "album":

		album_id, err := strconv.ParseInt(album, 10, 64)
		if err != nil {
			return []SongResponse{}, 40000, err
		}
		list, err := s.q.GetAlbumSongs(ctx, db.GetAlbumSongsParams{
			AlbumID: album_id,
			Limit:   pageSize,
			Offset:  (page - 1) * pageSize,
		})
		if err != nil {
			return []SongResponse{}, 40426, err
		}
		for _, element := range list {
			song, _ := s.GetSong(ctx, element)
			res = append(res, song)
		}
		return res, 200, nil

	case "singer":

		singer_id, err := strconv.ParseInt(singer, 10, 64)
		if err != nil {
			return []SongResponse{}, 40000, err
		}

		list, err := s.q.GetSingerSongs(ctx, db.GetSingerSongsParams{
			SingerID: singer_id,
			Limit:    pageSize,
			Offset:   (page - 1) * pageSize,
		})
		if err != nil {
			return []SongResponse{}, 40426, err
		}
		for _, element := range list {
			song, _ := s.GetSong(ctx, element)
			res = append(res, song)
		}
		return res, 200, nil

	case "genres":

		genre_id, err := strconv.ParseInt(genres, 10, 64)
		if err != nil {
			return []SongResponse{}, 40000, err
		}

		list, err := s.q.GetGenreSongs(ctx, db.GetGenreSongsParams{
			GenresID: genre_id,
			Limit:    pageSize,
			Offset:   (page - 1) * pageSize,
		})
		if err != nil {
			return []SongResponse{}, 40426, err
		}
		for _, element := range list {
			song, _ := s.GetSong(ctx, element)
			res = append(res, song)
		}
		return res, 200, nil

	case "":
		list, err := s.q.GetListSong(ctx, db.GetListSongParams{
			Limit:  pageSize,
			Offset: (page - 1) * pageSize,
		})
		if err != nil {
			return []SongResponse{}, 40426, err
		}
		for _, element := range list {
			song, _ := s.GetSong(ctx, element)
			res = append(res, song)
		}
		return res, 200, nil

	}
	return []SongResponse{}, 200, nil
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

func NewSongService(q *db.Queries, st *db.Store) ISongService {
	return &SongService{
		q:  q,
		st: st,
	}
}
