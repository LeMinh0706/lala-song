package singer

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type SingerService struct {
	q *db.Queries
}

// UpdateSinger implements ISingerService.
func (s *SingerService) UpdateSinger(ctx context.Context, id int64, fullname string, avt string) (*db.UpdateSingerRow, error) {
	update, err := s.q.UpdateSinger(ctx, db.UpdateSingerParams{
		ID:       id,
		Fullname: fullname,
		ImageUrl: avt,
	})
	if err != nil {
		return &db.UpdateSingerRow{}, err
	}
	return &update, nil
}

// DeleteSinger implements ISingerService.
func (s *SingerService) DeleteSinger(ctx context.Context, id int64) error {
	err := s.q.DeleteSinger(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetListSinger implements ISingerService.
func (s *SingerService) GetListSinger(ctx context.Context, page int32, page_size int32) ([]db.GetListSingerRow, int64) {
	list, _ := s.q.GetListSinger(ctx, db.GetListSingerParams{
		Limit:  page_size,
		Offset: (page - 1) * page_size,
	})
	count, _ := s.q.CountSinger(ctx)
	if list == nil {
		return []db.GetListSingerRow{}, count
	}
	return list, count
}

// GetSinger implements ISingerService.
func (s *SingerService) GetSinger(ctx context.Context, id int64) (*db.GetSingerRow, error) {
	singer, err := s.q.GetSinger(ctx, id)
	if err != nil {
		return &db.GetSingerRow{}, err
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
