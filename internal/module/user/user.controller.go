package user

import (
	"time"

	"github.com/LeMinh0706/lala-song/internal/handler"
	"github.com/LeMinh0706/lala-song/internal/middlewares"
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

// User godoc
// @Summary      Register user
// @Description  Join with us
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body Register true "request"
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /users/register [post]
func (u *UserController) Register(f *fiber.Ctx) error {
	var req Register
	if err := f.BodyParser(&req); err != nil {
		return handler.ValidateRegister(f, err)
	}
	_, err := u.service.Register(f.Context(), req)
	if err != nil {
		return handler.ValidateRegister(f, err)
	}
	return res.SuccessResponse(f, 201, nil)
}

// User godoc
// @Summary      Login user
// @Description  Login to be more handsome
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body Login true "request"
// @Success      200  {object}  LoginResponse
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /users/login [post]
func (u *UserController) Login(f *fiber.Ctx) error {
	var req Login

	if err := f.BodyParser(&req); err != nil {
		return res.ErrorResponse(f, 40000)
	}

	user, err := u.service.Login(f.Context(), req.Username, req.Password)
	if err != nil {
		return res.ErrorResponse(f, res.ErrLogin)
	}
	token, _ := u.token.CreateToken(user.Username, user.Name, time.Hour)

	response := LoginResponse{AccessToken: token}
	return res.SuccessResponse(f, 201, response)
}

// User godoc
// @Summary      It's you
// @Description  All your account is in here ->
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200  {object}  db.GetMeRow{}
// @Failure      500  {object}  res.ErrSwaggerJson
// @Router       /users/me [get]
func (u *UserController) GetMe(f *fiber.Ctx) error {
	auth := f.Locals(middlewares.AuthorizationPayloadKey).(*token.Payload)

	user, err := u.service.GetMe(f.Context(), auth.Username)
	if err != nil {
		return res.ErrorResponse(f, res.ErrUnauthorize)
	}

	return res.SuccessResponse(f, 200, user)
}
