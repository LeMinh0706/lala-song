package singer

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

type SingerController struct {
	service ISingerService
}

func NewSingerController(service ISingerService) *SingerController {
	return &SingerController{
		service: service,
	}
}

// Singer godoc
// @Summary      Create Singer
// @Description  Create Singer
// @Tags         Singers
// @Accept       multipart/form-data
// @Produce      json
// @Param        fullname formData string true "fullname"
// @Param        image formData file true "Image singer"
// @Security BearerAuth
// @Success      201  {object} 	db.Singer
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers [post]
func (s *SingerController) CreateSinger(f *fiber.Ctx) error {
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
	filename := fmt.Sprintf("upload/%s/%d%s", "singers", time.Now().Unix(), filepath.Ext(file.Filename))

	err = f.SaveFile(file, filename)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	fullname := f.FormValue("fullname")

	singer, err := s.service.CreateSinger(f.Context(), fullname, filename)

	if err != nil {
		return res.ErrorResponse(f, res.ErrAddSinger)
	}

	return res.SuccessResponse(f, 201, &singer)

}

// Singer godoc
// @Summary      Get singer with id
// @Description  Get singer with id
// @Tags         Singers
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Success      200  {object}  db.GetSingerRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers/{id} [get]
func (s *SingerController) GetSingerById(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}
	singer, err := s.service.GetSinger(f.Context(), id)
	if err != nil {
		return res.ErrorResponse(f, res.ErrSingerNotfound)
	}
	return res.SuccessResponse(f, 200, singer)
}

// Singer godoc
// @Summary      Get list singers
// @Description  Get list singers with page and page size (Limit-Offset)
// @Tags         Singers
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Success      200  {object}  SingersResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers [get]
func (s *SingerController) GetSingers(f *fiber.Ctx) error {
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}

	singer, total := s.service.GetListSinger(f.Context(), page, pageSize)

	return res.SuccessResponse(f, 200, SingersResponse{Singers: singer, Total: total})
}

// Singer godoc
// @Summary      Update singer
// @Description  Update singer
// @Tags         Singers
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Param        fullname formData string false "fullname"
// @Param        image formData file false "Image comment"
// @Security BearerAuth
// @Success      200  {object}  db.GetSingerRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers/{id} [put]
func (s *SingerController) UpdateSinger(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	fullname := f.FormValue("fullname")

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

		filename = fmt.Sprintf("upload/%s/%d%s", "singers", time.Now().Unix(), filepath.Ext(file.Filename))
		err = f.SaveFile(file, filename)
		if err != nil {
			return res.ErrorNonKnow(f, err.Error())
		}
	}

	singer, err := s.service.GetSinger(f.Context(), id)
	if err != nil {
		return res.ErrorResponse(f, res.ErrSingerNotfound)
	}

	if strings.TrimSpace(fullname) == "" {
		fullname = singer.Fullname
	}
	if filename == "" {
		filename = singer.ImageUrl
	}

	update, err := s.service.UpdateSinger(f.Context(), id, fullname, filename)

	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 201, update)
}

// Singer godoc
// @Summary      Delete singer with id
// @Description  Delete singer with id
// @Tags         Singers
// @Accept       json
// @Produce      json
// @Param        id path int true "ID"
// @Security BearerAuth
// @Success      204  "No content"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /singers/soft/{id} [post]
func (s *SingerController) DeleteSinger(f *fiber.Ctx) error {
	idStr := f.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequestId)
	}

	err = s.service.DeleteSinger(f.Context(), id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 204, nil)
}
