package handler

import (
	"github.com/LeMinh0706/lala-song/res"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRegister(f *fiber.Ctx, err error) error {
	if validate, ok := err.(validator.ValidationErrors); ok {
		for _, vali := range validate {
			switch vali.Tag() {
			case "min":
				if vali.Field() == "Username" {
					return res.ErrorResponse(f, 40008)
				} else if vali.Field() == "Gender" {
					return res.ErrorResponse(f, 40007)
				} else if vali.Field() == "Password" {
					return res.ErrorResponse(f, 40009)
				} else if vali.Field() == "Fullname" {
					return res.ErrorResponse(f, 40010)
				}
			case "max":
				if vali.Field() == "Username" {
					return res.ErrorResponse(f, 40008)
				} else if vali.Field() == "Gender" {
					return res.ErrorResponse(f, 40007)
				}
			case "required":
				if vali.Field() == "Username" {
					return res.ErrorResponse(f, 40008)
				} else if vali.Field() == "Gender" {
					return res.ErrorResponse(f, 40007)
				} else if vali.Field() == "Password" {
					return res.ErrorResponse(f, 40009)
				} else if vali.Field() == "Fullname" {
					return res.ErrorResponse(f, 40010)
				}
			}
		}
	}
	return res.ErrorResponse(f, 40000)
}
