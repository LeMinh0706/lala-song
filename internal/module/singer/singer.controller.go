package singer

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/LeMinh0706/lala-song/internal/handler"
	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/LeMinh0706/lala-song/util"
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
// @Param        fullname formData string true "fullname"
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
	const maxSize = 10 << 20
	if file.Size > maxSize {
		return res.ErrorResponse(f, 41300)
	}
	if !util.FileExtCheck(file.Filename) {
		return res.ErrorResponse(f, 40003)
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

	return res.SuccessResponse(f, 201, singer)

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
