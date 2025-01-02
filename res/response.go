package res

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(f *fiber.Ctx, code int, data interface{}) error {
	return f.Status(fiber.StatusOK).JSON(ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(f *fiber.Ctx, code int) error {
	return f.Status(fiber.StatusOK).JSON(ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

func ErrorNonKnow(f *fiber.Ctx, massage string) error {
	return f.Status(fiber.StatusOK).JSON(ResponseData{
		Code:    50000,
		Message: massage,
		Data:    nil,
	})
}

type ErrSwaggerJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
