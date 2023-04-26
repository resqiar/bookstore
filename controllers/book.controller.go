package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/outputs"

	"github.com/gofiber/fiber/v2"
)

func GetBookDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var book entities.Book

	// Find current logged in user data inside the database
	result := database.DB.First(&book, "id = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(&outputs.StatusOutput{
			Status: fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": fiber.StatusOK,
		"result": book,
	})
}
