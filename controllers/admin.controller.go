package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/inputs"
	"bookstore/libs"
	"bookstore/outputs"

	"github.com/gofiber/fiber/v2"
)

func AddBook(c *fiber.Ctx) error {
	var payload inputs.AddBookInput

	// Bind request body to payload struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Validate the struct to follow the format of
	// the defined struct. see inputs/add-book.input.go
	err := libs.InputValidator(payload)
	if err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusBadRequest,
			Message: err,
		})
	}

	newBook := entities.Book{
		Title:       payload.Title,
		Description: payload.Description,
		ImageURL:    payload.ImageURL,
		ReleaseDate: payload.ReleaseDate,
		Author:      payload.Author,
		Price:       payload.Price,
	}

	result := database.DB.Create(&newBook)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusInternalServerError,
			Message: result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
	})
}

func GetBooks(c *fiber.Ctx) error {
	var books []entities.Book

	// Find current logged in user data inside the database
	result := database.DB.Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": books,
	})
}
