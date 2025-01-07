package search

import (
	"context"
	"database/sql"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/song"
)

type SearchService struct {
	q    *db.Queries
	song song.ISongService
}

// SearchLyric implements ISearchService.
func (s *SearchService) SearchLyric(ctx context.Context, lyric string, page, pageSize int32) ([]song.SongResponse, error) {
	var res []song.SongResponse
	list, err := s.q.SearchSongsByLyrics(ctx, db.SearchSongsByLyricsParams{
		PlaintoTsquery: lyric,
		Limit:          pageSize,
		Offset:         (page - 1) * pageSize,
	})
	if err != nil {
		return []song.SongResponse{}, err
	}

	for _, element := range list {
		play, _ := s.song.GetSong(ctx, element)
		res = append(res, play)
	}
	return res, nil
}

// SearchSong implements ISearchService.
func (s *SearchService) SearchSong(ctx context.Context, name string, page, pageSize int32) ([]song.SongResponse, error) {
	var res []song.SongResponse
	list, err := s.q.SearchSong(ctx, db.SearchSongParams{
		Column1: sql.NullString{String: name, Valid: true},
		Limit:   pageSize,
		Offset:  (page - 1) * pageSize,
	})
	if err != nil {
		return []song.SongResponse{}, err
	}
	if len(list) == 0 {
		return []song.SongResponse{}, nil
	}

	for _, element := range list {
		play, _ := s.song.GetSong(ctx, element)
		res = append(res, play)
	}
	return res, nil
}

func NewSearchService(q *db.Queries, song song.ISongService) ISearchService {
	return &SearchService{
		q:    q,
		song: song,
	}
}
