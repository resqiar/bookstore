package controllers

import (
	"bookstore/database"
	"bookstore/entities"
	"bookstore/inputs"
	"bookstore/libs"
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
	result = database.DB.Preload("Book").Find(&cart, "user_id = ?", user.ID)
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

func AddToCart(c *fiber.Ctx) error {
	var payload inputs.AddToCartInput

	// Bind request body to payload struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// Validate the struct to follow the format of
	// the defined struct. see inputs/cart.input.go
	err := libs.InputValidator(payload)
	if err != "" {
		return c.Status(fiber.StatusBadRequest).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusBadRequest,
			Message: err,
		})
	}

	var isExist entities.Cart
	result := database.DB.First(&isExist, "user_id = ? AND book_id = ?", payload.UserID, payload.BookID)
	if result.Error != nil && result.Error.Error() == "record not found" {
		newCart := entities.Cart{
			UserID:   payload.UserID,
			BookID:   payload.BookID,
			Quantity: payload.Quantity,
		}

		result = database.DB.Create(&newCart)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&outputs.ErrorOutput{
				Status:  fiber.StatusInternalServerError,
				Message: result.Error.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&outputs.StatusOutput{
			Status: fiber.StatusOK,
		})
	}

	// AT THIS POINT THE CART IS ITEM IS EXIST
	// WE JUST NEED TO UPDATE THE QUANTITY
	isExist.Quantity += 1

	// Save the updated data to the database
	if err := database.DB.Save(&isExist).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&outputs.ErrorOutput{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&outputs.StatusOutput{
		Status: fiber.StatusOK,
	})
}
