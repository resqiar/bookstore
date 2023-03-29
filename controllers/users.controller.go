package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/outputs"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var user entities.User

	// Find current logged in user data inside the database
	result := database.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"user":   user,
	})
}
