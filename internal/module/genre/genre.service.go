package genre

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type IGenreService interface {
	CreateGenre(ctx context.Context, name, image_url string) (*db.Genre, error)
	GetGenreById(ctx context.Context, id int64) (*db.Genre, error)
	GetListGenres(ctx context.Context, page, pageSize int32) ([]db.Genre, int64)
	UpdateGenre(ctx context.Context, id int64, name, image_url string) (*db.Genre, error)
	DeleteGenre(ctx context.Context, id int64) error
}
