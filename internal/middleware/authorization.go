package middleware

import "github.com/gofiber/fiber/v2"

func (m *MiddleWare) Authorization(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin")

	if isAdmin == false {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "admin only",
		})
	}

	return ctx.Next()
}