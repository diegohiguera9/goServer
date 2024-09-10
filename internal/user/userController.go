package user

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("I am creating a user...")
}
