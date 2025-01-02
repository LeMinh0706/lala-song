package singer

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

type SingerController struct {
	service ISingerService
	token   token.Maker
}

func NewSingerController(service ISingerService, token token.Maker) *SingerController {
	return &SingerController{
		service: service,
		token:   token,
	}
}

// Singer godoc
// @Summary      Create Singer
// @Description  Create Singer
// @Tags         Singers
// @Accept       multipart/form-data
// @Produce      json
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      201
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers [post]
func (s *SingerController) CreateSinger(f *fiber.Ctx) error {
	file, err := f.FormFile("image")
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestMime)
	}
	filename := fmt.Sprintf("upload/%s/%d%s", "genres", time.Now().Unix(), filepath.Ext(file.Filename))
	err = f.SaveFile(file, filename)
	return res.SuccessResponse(f, 201, nil)
}
