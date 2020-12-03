package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// UserList returns a lits of users
func UserList(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"user":    "testUser",
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"user":    "testUser",
	})
}
