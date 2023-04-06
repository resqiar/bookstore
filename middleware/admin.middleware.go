package middleware

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/outputs"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	var user entities.User
	result := database.DB.First(&user, "id = ?", userId)
	fmt.Println(userId)
	if userId == nil || result.Error != nil || !user.IsAdmin {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	return c.Next()
}
