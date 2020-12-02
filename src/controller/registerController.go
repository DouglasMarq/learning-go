package controller

import "github.com/gofiber/fiber/v2"

// RegisterController test
func RegisterController(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "OK",
	})
}

