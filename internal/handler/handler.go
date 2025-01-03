package handler

import (
	"strconv"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRegister(f *fiber.Ctx, err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errorMap := map[string]map[string]int{
			"Username": {
				"min":      40008,
				"max":      40008,
				"required": 40008,
			},
			"Password": {
				"min":      40009,
				"required": 40009,
			},
			"Fullname": {
				"min":      40010,
				"required": 40010,
			},
			"Gender": {
				"min":      40007,
				"max":      40007,
				"required": 40007,
			},
		}

		for _, vali := range validationErrors {
			if fieldErrors, ok := errorMap[vali.Field()]; ok {
				if errorCode, ok := fieldErrors[vali.Tag()]; ok {
					return res.ErrorResponse(f, errorCode)
				}
			}
		}
	}

	return res.ErrorNonKnow(f, err.Error())
}

func CheckQuery(f *fiber.Ctx, pageStr, pageSizeStr string) (int32, int32, int) {
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		return 0, 0, 40001
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		return 0, 0, 40002
	}
	if page <= 0 || pageSize <= 0 {
		return 0, 0, 40002
	}
	return int32(page), int32(pageSize), 0
}
