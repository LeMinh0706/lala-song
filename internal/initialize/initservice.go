package initialize

import (
	"database/sql"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/internal/module/album"
	"github.com/LeMinh0706/lala-song/internal/module/genre"
	"github.com/LeMinh0706/lala-song/internal/module/singer"
	"github.com/LeMinh0706/lala-song/internal/module/song"
	"github.com/LeMinh0706/lala-song/internal/module/user"
	"github.com/LeMinh0706/lala-song/util"
)

type ISevice struct {
	UserService   user.IUserService
	SingerService singer.ISingerService
	GenreService  genre.IGenreService
	AlbumService  album.IAlbumService
	SongService   song.ISongService
}

func InitService(pg *sql.DB, config util.Config) *ISevice {
	store := db.NewStore(pg)

	q := db.New(pg)

	userService := user.NewUserService(q)
	singerService := singer.NewSingerService(q)
	genreService := genre.NewGenreService(q)
	albumService := album.NewAlbumService(q, singerService)
	songService := song.NewSongService(q, store)

	return &ISevice{
		UserService:   userService,
		SingerService: singerService,
		GenreService:  genreService,
		AlbumService:  albumService,
		SongService:   songService,
	}
}
