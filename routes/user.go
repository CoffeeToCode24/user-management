package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/user-managemnet/database"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(database.Users)
}
