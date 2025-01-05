package genre

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/gofiber/fiber/v2"
)

type GenreController struct {
	service IGenreService
}

func NewGenreController(service IGenreService) *GenreController {
	return &GenreController{
		service: service,
	}
}

// Genres godoc
// @Summary      Create Genre
// @Description  Create Genre
// @Tags         Genres
// @Accept       multipart/form-data
// @Produce      json
// @Param        name formData string true "name"
// @Param        image formData file true "Image genre"
// @Security BearerAuth
// @Success      201  {object} 	db.Genre
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /genres [post]
func (g *GenreController) CreateGenre(f *fiber.Ctx) error {
	file, err := f.FormFile("image")
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	const maxSize = 10 << 20
	if file.Size > maxSize {
		return res.ErrorResponse(f, 41300)
	}
	if file != nil {
		if !util.FileExtCheck(file.Filename) {
			return res.ErrorResponse(f, res.ErrBadRequestMime)
		}
	}
	filename := fmt.Sprintf("upload/%s/%d%s", "genres", time.Now().Unix(), filepath.Ext(file.Filename))

	err = f.SaveFile(file, filename)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	name := f.FormValue("name")

	genre, err := g.service.CreateGenre(f.Context(), name, filename)

	if err != nil {
		return res.ErrorResponse(f, res.ErrAddGenre)
	}

	return res.SuccessResponse(f, 201, &genre)

}
