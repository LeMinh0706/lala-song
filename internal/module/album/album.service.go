package album

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type IAlbumService interface {
	CreateAlbum(ctx context.Context, singer_id int64, name, image_url string) (*db.CreateAlbumRow, int)
	GetAlbumById(ctx context.Context, id int64) (*db.GetAlbumRow, error)
	GetListAlbum(ctx context.Context, page, pageSize int32) ([]GetAlbumResponse, error)
	GetSingerAlbun(ctx context.Context, singer_id int64, page, pageSize int32) ([]GetAlbumResponse, error)
	UpdateAlbum(ctx context.Context, id int64, name, image_url string) (*db.UpdateAlbumParams, error)
	DeleteAlbum(ctx context.Context, id int64) error
}
