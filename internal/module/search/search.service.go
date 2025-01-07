package search

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/module/song"
)

type ISearchService interface {
	SearchSong(ctx context.Context, name string, page, pageSize int32) ([]song.SongResponse, error)
	SearchLyric(ctx context.Context, lyric string, page, pageSize int32) ([]song.SongResponse, error)
}
