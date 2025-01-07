package favorite

import (
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

func (fc *FavoriteController) CreateLikeSong(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)
	song := f.Params("id")

	id, err := uuid.Parse(song)
	if err != nil {
		return res.ErrorResponse(f, res.ErrBadRequest)
	}

	create, err := fc.service.CreateLikeSong(f.Context(), auth.Username, id)
	if err != nil {
		return res.ErrorNonKnow(f, err.Error())
	}

	return res.SuccessResponse(f, 201, create)
}

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
