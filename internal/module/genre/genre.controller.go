package genre

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

// Genre godoc
// @Summary      Get genre with id
// @Description  Get genre with id
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200  {object}  db.Genre
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /genres/{id} [get]
func (g *GenreController) GetGenreById(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}
	genre, err := g.service.GetGenreById(f.Context(), id)
	if err != nil {
		return res.ErrorResponse(f, res.ErrGenreNotFound)
	}
	return res.SuccessResponse(f, 200, genre)
}

// Genre godoc
// @Summary      Get list genres
// @Description  Get list genres with page and page size (Limit-Offset)
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      200  {object}  GenresResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /genres [get]
func (g *GenreController) GetGenres(f *fiber.Ctx) error {
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}

	genre, total := g.service.GetListGenres(f.Context(), page, pageSize)

	return res.SuccessResponse(f, 200, GenresResponse{Genres: genre, Total: total})
}

// Genre godoc
// @Summary      Update genre
// @Description  Update genre
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        name formData string false "name"
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      200  {object}  db.Genre
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /genres/{id} [put]
func (g *GenreController) UpdateGenre(f *fiber.Ctx) error {
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

		filename = fmt.Sprintf("upload/%s/%d%s", "genres", time.Now().Unix(), filepath.Ext(file.Filename))
		err = f.SaveFile(file, filename)
		if err != nil {
			return res.ErrorNonKnow(f, err.Error())
		}
	}

	genre, err := g.service.GetGenreById(f.Context(), id)
	if err != nil {
		return res.ErrorResponse(f, res.ErrGenreNotFound)
	}

	if strings.TrimSpace(name) == "" {
		name = genre.Name
	}
	if filename == "" {
		filename = genre.ImageUrl
	}

	update, err := g.service.UpdateGenre(f.Context(), id, name, filename)

	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 201, update)
}

// Genre godoc
// @Summary      Delete genre with id
// @Description  Delete genre with id
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /genres/{id} [delete]
func (g *GenreController) DeleteGenre(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	err = g.service.DeleteGenre(f.Context(), id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 204, nil)
}
