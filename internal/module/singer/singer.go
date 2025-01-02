package singer

import "github.com/LeMinh0706/lala-song/internal/db"

type SingersResponse struct {
	Singers []db.GetListSingerRow `json:"singers"`
	Total   int64                 `json:"total"`
}
