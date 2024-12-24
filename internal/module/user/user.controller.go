package user

import (
	"github.com/LeMinh0706/lala-song/res"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service IUserService
}

func NewUserController(service IUserService) *UserController {
	return &UserController{
		service: service,
	}
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
	return res.SuccessResponse(f, 201, user)
}
