package server

import (
	"github.com/LeMinh0706/lala-song/internal/initialize"
	"github.com/LeMinh0706/lala-song/internal/module/album"
	"github.com/LeMinh0706/lala-song/internal/module/favorite"
	"github.com/LeMinh0706/lala-song/internal/module/genre"
	"github.com/LeMinh0706/lala-song/internal/module/singer"
	"github.com/LeMinh0706/lala-song/internal/module/song"
	"github.com/LeMinh0706/lala-song/internal/module/user"
	_ "github.com/LeMinh0706/lala-song/swag/docs"
	"github.com/gofiber/swagger"
)

func (s *Server) NewRouter() {

	initService := initialize.InitService(s.DBConn, s.Config)

	s.Router.Get("/swagger/*", swagger.HandlerDefault)
	a := s.Router.Group("/api")
	{
		user.NewUserRouter(a, initService.UserService, s.TokenMaker)
		singer.NewSingerRouter(a, initService.SingerService, s.TokenMaker)
		genre.NewGenreRouter(a, initService.GenreService, s.TokenMaker)
		album.NewAlbumRouter(a, initService.AlbumService, s.TokenMaker)
		song.NewSongRouter(a, initService.SongService, s.TokenMaker)
		favorite.NewFavoriteRouter(a, initService.FavoriteService, s.TokenMaker)
	}

}
