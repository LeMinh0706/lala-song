package favorite

import (
	"github.com/LeMinh0706/lala-song/internal/handler"
	"github.com/LeMinh0706/lala-song/internal/middlewares"
	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FavoriteController struct {
	service IFavoriteService
	token   token.Maker
}

func NewFavoriteController(service IFavoriteService, token token.Maker) *FavoriteController {
	return &FavoriteController{
		service: service,
		token:   token,
	}
}

// Favorite godoc
// @Summary      Like song
// @Description  Like song
// @Tags         Favorite
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Security BearerAuth
// @Success      200  {object}  db.Favorite
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /favorite/{id} [post]
func (fc *FavoriteController) CreateLikeSong(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)
	song := f.Params("id")

	id, err := uuid.Parse(song)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequest)
	}

	_, err = fc.service.CreateLikeSong(f.Context(), auth.Username, id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 201, nil)
}

// Song godoc
// @Summary      Get like
// @Description  Xem thử có like bài hát hay chưa
// @Tags         Favorite
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Security BearerAuth
// @Success      200  {object}  db.GetSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /favorite/{id} [get]
func (fc *FavoriteController) GetLikeSong(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)

	song := f.Params("id")

	id, err := uuid.Parse(song)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequest)
	}

	like, _ := fc.service.GetLikeSong(f.Context(), auth.Username, id)

	return res.SuccessResponse(f, 201, like)
}

// Song godoc
// @Summary      Get list favorite
// @Description  Get list favorite with page and page size (Limit-Offset)
// @Tags         Favorite
// @Accept       json
// @Produce      json
// @Param        page query int true "Page"
// @Param        page_size query int true "Page Size"
// @Security BearerAuth
// @Success      200  {object}  []db.GetSongRow
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /favorite [get]
func (fc *FavoriteController) GetListSongs(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)
	pageStr := f.Query("page")
	pageSizeStr := f.Query("page_size")

	page, pageSize, numErr := handler.CheckQuery(f, pageStr, pageSizeStr)
	if numErr != 0 {
		return res.ErrorResponse(f, numErr)
	}

	list, err := fc.service.GetListSong(f.Context(), auth.Username, page, pageSize)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 200, list)
}

// Favorite godoc
// @Summary      Like song
// @Description  Like song
// @Tags         Favorite
// @Accept       json
// @Produce      json
// @Param        id path string true "ID"
// @Security BearerAuth
// @Success      204 "No content"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /favorite/{id} [delete]
func (fc *FavoriteController) UnlikeSong(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)

	song := f.Params("id")

	id, err := uuid.Parse(song)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequest)
	}

	err = fc.service.Unlike(f.Context(), auth.Username, id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}
	return res.SuccessResponse(f, 204, nil)
}
