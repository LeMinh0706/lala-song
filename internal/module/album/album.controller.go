package album

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/gofiber/fiber/v2"
)

type AlbumController struct {
	service IAlbumService
}

func NewAlBumController(service IAlbumService) *AlbumController {
	return &AlbumController{
		service: service,
	}
}

// Albums godoc
// @Summary      Create Album
// @Description  Create Album
// @Tags         Albums
// @Accept       multipart/form-data
// @Produce      json
// @Param        singer_id formData int true "Singer Id"
// @Param        name formData string true "name"
// @Param        image formData file true "Image genre"
// @Security BearerAuth
// @Success      201  {object} 	db.Album
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /albums [post]
func (a *AlbumController) CreateAlbum(f *fiber.Ctx) error {
	singerIdStr := f.FormValue("singer_id")
	singer_id, err := strconv.ParseInt(singerIdStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

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
	filename := fmt.Sprintf("upload/%s/%d%s", "album", time.Now().Unix(), filepath.Ext(file.Filename))

	err = f.SaveFile(file, filename)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	name := f.FormValue("name")

	album, code := a.service.CreateAlbum(f.Context(), singer_id, name, filename)

	if code != 0 {
		return res.ErrorResponse(f, code)
	}
	return res.SuccessResponse(f, 201, album)
}
