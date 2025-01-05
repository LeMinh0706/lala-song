package genre

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type GenreService struct {
	q *db.Queries
}

// CreateGenre implements IGenreService.
func (g *GenreService) CreateGenre(ctx context.Context, name string, image_url string) (*db.Genre, error) {
	genre, err := g.q.CreateGenre(ctx, db.CreateGenreParams{
		Name:     name,
		ImageUrl: image_url,
	})
	if err != nil {
		return &db.Genre{}, err
	}
	return &genre, nil
}

// DeleteGenre implements IGenreService.
func (g *GenreService) DeleteGenre(ctx context.Context, id int64) error {
	err := g.q.DeleteGenreSong(ctx, id)
	if err != nil {
		return err
	}
	err = g.q.DeleteGenre(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetGenreById implements IGenreService.
func (g *GenreService) GetGenreById(ctx context.Context, id int64) (*db.Genre, error) {
	genre, err := g.q.GetGenre(ctx, id)
	if err != nil {
		return &db.Genre{}, err
	}
	return &genre, nil
}

// GetListGenres implements IGenreService.
func (g *GenreService) GetListGenres(ctx context.Context, page int32, pageSize int32) ([]db.Genre, int64) {
	list, _ := g.q.GetListGenre(ctx, db.GetListGenreParams{Limit: pageSize, Offset: (page - 1) * pageSize})
	count, _ := g.q.CountGenre(ctx)
	if list == nil {
		return []db.Genre{}, count
	}
	return list, count
}

// UpdateGenre implements IGenreService.
func (g *GenreService) UpdateGenre(ctx context.Context, id int64, name string, image_url string) (*db.Genre, error) {
	update, err := g.q.UpdateGenre(ctx, db.UpdateGenreParams{
		ID:       id,
		Name:     name,
		ImageUrl: image_url,
	})
	if err != nil {
		return &db.Genre{}, err
	}
	return &update, err
}

func NewGenreService(q *db.Queries) IGenreService {
	return &GenreService{
		q: q,
	}
}
