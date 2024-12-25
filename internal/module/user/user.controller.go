package user

import (
	"time"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service IUserService
	token   token.Maker
}

func NewUserController(service IUserService, token token.Maker) *UserController {
	return &UserController{
		service: service,
		token:   token,
	}
}

func (u *UserController) Register(f *fiber.Ctx) error {
	var req Register
	if err := f.BodyParser(&req); err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}
	_, err := u.service.Register(f.Context(), req)
	if err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}
	token, err := u.token.CreateToken(req.Username, time.Hour)
	if err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}
	return res.SuccessResponse(f, 201, LoginResponse{AccessToken: token})
}

func (u *UserController) Login(f *fiber.Ctx) error {
	var req Login

	if err := f.BodyParser(&req); err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}

	user, err := u.service.Login(f.Context(), req.Username, req.Password)
	if err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}
	token, err := u.token.CreateToken(user.Username, time.Hour)
	if err != nil {
		return res.ErrorNonKnow(f, 50000, err.Error())
	}
	response := LoginResponse{AccessToken: token}
	return res.SuccessResponse(f, 201, response)
}
