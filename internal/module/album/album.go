package album

import "github.com/LeMinh0706/lala-song/internal/db"

type GetAlbumResponse struct {
	Album []db.GetAlbumRow `json:"album"`
	Total int64            `json:"total"`
}
