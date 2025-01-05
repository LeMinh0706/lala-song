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
	err := a.q.DeleteAlbum(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetAlbumById implements IAlbumService.
func (a *AlbumService) GetAlbumById(ctx context.Context, id int64) (*db.GetAlbumRow, error) {
	album, err := a.q.GetAlbum(ctx, id)
	if err != nil {
		return &db.GetAlbumRow{}, err
	}
	return &album, nil
}

// GetListAlbum implements IAlbumService.
func (a *AlbumService) GetListAlbum(ctx context.Context, page int32, pageSize int32) ([]db.GetAlbumRow, int64) {
	list, _ := a.q.GetListAlbum(ctx, db.GetListAlbumParams{
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})

	count, _ := a.q.CountAlbum(ctx)

	if len(list) == 0 {
		return []db.GetAlbumRow{}, count
	}

	var albums []db.GetAlbumRow

	for _, element := range list {
		album, _ := a.q.GetAlbum(ctx, element)
		albums = append(albums, album)
	}
	return albums, count
}

// GetSingerAlbum implements IAlbumService.
func (a *AlbumService) GetSingerAlbum(ctx context.Context, singer_id int64, page int32, pageSize int32) ([]db.GetAlbumRow, int64, error) {
	var albums []db.GetAlbumRow
	_, err := a.q.GetSinger(ctx, singer_id)
	if err != nil {
		return []db.GetAlbumRow{}, 0, err
	}

	list, _ := a.q.GetSingerAlbums(ctx, db.GetSingerAlbumsParams{
		Limit:    pageSize,
		Offset:   (page - 1) * pageSize,
		SingerID: singer_id,
	})
	total, _ := a.q.CountSingerAlbum(ctx, singer_id)

	if list == nil {
		return []db.GetAlbumRow{}, total, nil
	}

	for _, element := range list {
		album, _ := a.q.GetAlbum(ctx, element)
		albums = append(albums, album)
	}
	return albums, total, nil
}

// UpdateAlbum implements IAlbumService.
func (a *AlbumService) UpdateAlbum(ctx context.Context, id int64, name string, image_url string) (*db.UpdateAlbumRow, error) {
	update, err := a.q.UpdateAlbum(ctx, db.UpdateAlbumParams{ID: id, Name: name, ImageUrl: image_url})
	if err != nil {
		return &db.UpdateAlbumRow{}, err
	}
	return &update, nil

}

func NewAlbumService(q *db.Queries, s singer.ISingerService) IAlbumService {
	return &AlbumService{
		q: q,
		s: s,
	}
}
