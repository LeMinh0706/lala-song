package song

import "github.com/LeMinh0706/lala-song/internal/db"

type SongResponse struct {
	Genres []db.Genre                 `json:"genres"`
	Song   db.GetSongRow              `json:"song"`
	Singer []db.GetSingersWithSongRow `json:"singers"`
}
