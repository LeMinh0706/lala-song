package singer

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type SingerService struct {
	q *db.Queries
}

// GetListSinger implements ISingerService.
func (s *SingerService) GetListSinger(ctx context.Context, page int32, page_size int32) ([]db.Singer, error) {
	panic("unimplemented")
}

// GetSinger implements ISingerService.
func (s *SingerService) GetSinger(ctx context.Context, id int64) (*db.Singer, error) {
	singer, err := s.q.GetSinger(ctx, id)
	if err != nil {
		return &db.Singer{}, nil
	}
	return &singer, err
}

// CreateSinger implements ISingerService.
func (s *SingerService) CreateSinger(ctx context.Context, fullname string, avatar string) (*db.Singer, error) {
	create, err := s.q.CreateSinger(ctx, db.CreateSingerParams{
		Fullname: fullname,
		ImageUrl: avatar,
	})
	if err != nil {
		return &db.Singer{}, err
	}
	return &create, nil
}

func NewSingerService(q *db.Queries) ISingerService {
	return &SingerService{
		q: q,
	}
}
