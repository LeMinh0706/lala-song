package middlewares

import (
	"strings"

	"github.com/LeMinh0706/lala-song/res"
	"github.com/LeMinh0706/lala-song/token"
	"github.com/gofiber/fiber/v2"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthorizeMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			return res.ErrorResponse(ctx, 40101)
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer") {
			authorizationHeader = "Bearer " + authorizationHeader
			ctx.Set(AuthorizationHeaderKey, authorizationHeader)
		}

		fields := strings.Fields(authorizationHeader)

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return res.ErrorResponse(ctx, 40101)
		}
		ctx.Locals(AuthorizationPayloadKey, payload)
		return ctx.Next()
	}
}
func AuthorizeAdminMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			return res.ErrorResponse(ctx, 40101)
		}

		if !strings.HasPrefix(authorizationHeader, "Bearer") {
			authorizationHeader = "Bearer " + authorizationHeader
			ctx.Set(AuthorizationHeaderKey, authorizationHeader)
		}

		fields := strings.Fields(authorizationHeader)

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return res.ErrorResponse(ctx, 40101)
		}
		// Kiểm tra role
		if payload.Rolename != "Admin" {
			return res.ErrorResponse(ctx, 40301)
		}
		ctx.Locals(AuthorizationPayloadKey, payload)

		return ctx.Next()
	}
}
