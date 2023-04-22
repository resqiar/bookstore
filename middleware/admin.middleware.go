package middleware

import (
	"bookstore/database"
	"bookstore/entities"

	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	var user entities.User
	result := database.DB.First(&user, "id = ?", userId)

	if userId == nil || result.Error != nil || !user.IsAdmin {
		return c.Status(fiber.StatusTemporaryRedirect).Redirect("/login")
	}

	return c.Next()
}
