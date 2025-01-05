package genre

import "github.com/LeMinh0706/lala-song/internal/db"

type GenresResponse struct {
	Genres []db.Genre `json:"genres"`
	Total  int64      `json:"total"`
}
