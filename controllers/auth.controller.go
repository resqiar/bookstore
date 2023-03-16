package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/inputs"
	"bookstore/libs"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var payload inputs.LoginInput
	// Bind request body to payload struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error()})
	}

	// Validate the struct to follow the format of
	// the defined struct. see inputs/login.input.go
	err := libs.AuthValidator(payload)
	if err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err,
		})
	}

	// Get username & password from database
	var user entities.User
	result := database.DB.First(&user, "username = ?", payload.Username)

	// If query result not found, return 400
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Username or password is not correct",
		})
	}

	// Match the password given with the hashed password
	matched := libs.ComparePassword(user.Password, payload.Password)
	if !matched {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Username or password is not correct",
		})
	}

	// If Match -> generate JWT
	token, tokenizationErr := libs.GenerateToken(user.ID)
	if tokenizationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": tokenizationErr.Error(),
		})
	}

	// and send back JWT
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"token":  token,
	})
}
