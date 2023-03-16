package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/inputs"
	"bookstore/libs"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var payload inputs.RegisterInput

	// Bind request body to payload struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error()})
	}

	// Validate the struct to follow the format of
	// the defined struct. see inputs/register.input.go
	err := libs.AuthValidator(payload)
	if err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err,
		})
	}

	// Hash password before saved to the database
	hashedPassword, hashErr := libs.HashPassword(payload.Password)
	if hashErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": hashErr,
		})
	}

	// update payload password to the hashed version
	payload.Password = hashedPassword

	// Create a new user with the specified username & password
	newUser := entities.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	// Save the created input to the database
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": result.Error.Error(),
		})
	}

	token, tokenizationErr := libs.GenerateToken(newUser.ID)
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
