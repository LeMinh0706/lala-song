package singer

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
)

type ISingerService interface {
	CreateSinger(ctx context.Context, fullname, avatar string) (*db.Singer, error)
	GetSinger(ctx context.Context, id int64) (*db.GetSingerRow, error)
	GetListSinger(ctx context.Context, page, page_size int32) ([]db.GetListSingerRow, int64)
}
