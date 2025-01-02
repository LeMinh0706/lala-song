//go:build wireinject

package wire

import (
	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/singer"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/google/wire"
)

func InitSingerRouterHandler(q *db.Queries, token token.Maker) (*singer.SingerController, error) {
	wire.Build(
		singer.NewSingerService,
		singer.NewSingerController,
	)
	return new(singer.SingerController), nil
}
