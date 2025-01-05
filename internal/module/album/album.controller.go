package album

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/LeMinh0706/lala-song/internal/handler"
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
// @Success      201  {object} 	db.CreateAlbumRow
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

// Album godoc
// @Summary      Get list albums
// @Description  Get list albums with page and page size (Limit-Offset)
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        singer_id query int false "Singer ID"
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      200  {object}  GetAlbumResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /albums [get]
func (a *AlbumController) GetAlbums(f *fiber.Ctx) error {
	idStr := f.Query("singer_id")
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}

	if idStr == "" {
		albums, count := a.service.GetListAlbum(f.Context(), page, pageSize)
		return res.SuccessResponse(f, 200, GetAlbumResponse{Album: albums, Total: count})
	}

	singer_id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	album, count, err := a.service.GetSingerAlbum(f.Context(), singer_id, page, pageSize)
	if err != nil {
		return res.ErrorResponse(f, res.ErrSingerNotfound)
	}

	return res.SuccessResponse(f, 200, GetAlbumResponse{Album: album, Total: count})
}

// Album godoc
// @Summary      Update album
// @Description  Update album
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        name formData string false "name"
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      200  {object}  db.CreateAlbumRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /albums/{id} [put]
func (a *AlbumController) UpdateAlbum(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	name := f.FormValue("name")

	var filename string
	file, err := f.FormFile("image")
	if err == nil && file != nil {
		const maxSize = 10 << 20
		if file.Size > maxSize {
			return res.ErrorResponse(f, 41300)
		}
		if !util.FileExtCheck(file.Filename) {
			return res.ErrorResponse(f, res.ErrBadRequestMime)
		}

		filename = fmt.Sprintf("upload/%s/%d%s", "album", time.Now().Unix(), filepath.Ext(file.Filename))
		err = f.SaveFile(file, filename)
		if err != nil {
			return res.ErrorNonKnow(f, err.Error())
		}
	}

	album, err := a.service.GetAlbumById(f.Context(), id)
	if err != nil {
		return res.ErrorResponse(f, res.ErrAlbumNotFound)
	}

	if strings.TrimSpace(name) == "" {
		name = album.Name
	}
	if filename == "" {
		filename = album.ImageUrl
	}

	update, err := a.service.UpdateAlbum(f.Context(), id, name, filename)

	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 201, update)
}

// Album godoc
// @Summary      Delete album with id
// @Description  Delete album with id
// @Tags         Albums
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /album/soft/{id} [post]
func (a *AlbumController) DeleteAlbum(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	err = a.service.DeleteAlbum(f.Context(), id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 204, nil)
}
