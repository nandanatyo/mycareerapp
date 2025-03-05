package middleware

import (
	"mycareerapp/internal/infra/jwt"

	"github.com/gofiber/fiber/v2"
)

type MiddleWareI interface {
	Authentication(ctx *fiber.Ctx) error
	Authorization(ctx *fiber.Ctx) error 
}

type MiddleWare struct {
	jwt jwt.JWTI
}

func NewMiddleware(jwt jwt.JWTI) MiddleWareI {
	return &MiddleWare{
		jwt: jwt,
	}
}
