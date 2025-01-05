package album

import (
	"context"
	"sync"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/singer"
)

type AlbumService struct {
	q *db.Queries
	s singer.ISingerService
}

// CreateAlbum implements IAlbumService.
func (a *AlbumService) CreateAlbum(ctx context.Context, singer_id int64, name string, image_url string) (*db.CreateAlbumRow, int) {
	var res *db.CreateAlbumRow
	var wg sync.WaitGroup
	errChan := make(chan int, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := a.q.GetSinger(ctx, singer_id)
		if err != nil {
			errChan <- 40423
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		album, err := a.q.CreateAlbum(ctx, db.CreateAlbumParams{
			Name:     name,
			ImageUrl: image_url,
			SingerID: singer_id,
		})
		if err != nil {
			errChan <- 40018
		}
		res = &album
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		return nil, err
	}

	return res, 0
}

// DeleteAlbum implements IAlbumService.
func (a *AlbumService) DeleteAlbum(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetAlbumById implements IAlbumService.
func (a *AlbumService) GetAlbumById(ctx context.Context, id int64) (*db.GetAlbumRow, error) {
	panic("unimplemented")
}

// GetListAlbum implements IAlbumService.
func (a *AlbumService) GetListAlbum(ctx context.Context, page int32, pageSize int32) ([]GetAlbumResponse, error) {
	panic("unimplemented")
}

// GetSingerAlbun implements IAlbumService.
func (a *AlbumService) GetSingerAlbun(ctx context.Context, singer_id int64, page int32, pageSize int32) ([]GetAlbumResponse, error) {
	panic("unimplemented")
}

// UpdateAlbum implements IAlbumService.
func (a *AlbumService) UpdateAlbum(ctx context.Context, id int64, name string, image_url string) (*db.UpdateAlbumParams, error) {
	panic("unimplemented")
}

func NewAlbumService(q *db.Queries, s singer.ISingerService) IAlbumService {
	return &AlbumService{
		q: q,
		s: s,
	}
}
