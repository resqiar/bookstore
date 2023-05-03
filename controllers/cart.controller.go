package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/outputs"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUserCart(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	var user entities.User

	// Find current logged in user data inside the database
	result := database.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	// Find the data in Cart table where the user id == current user id
	var cart []entities.Cart
	result = database.DB.Find(&cart, "user_id = ?", user.ID)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusInternalServerError,
			Message: result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": fiber.StatusOK,
		"result": cart,
	})
}
