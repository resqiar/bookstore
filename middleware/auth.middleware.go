package middleware

import (
	"bookstore/config"
	"bookstore/libs"
	"bookstore/outputs"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	// token extracted from Cookie
	token := c.Cookies("token")

	// if there is no token inside the Cookie,
	// send 401 status of unauthorized
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	claims, valid := libs.ParseToken(token)
	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(&outputs.StatusOutput{
			Status: fiber.StatusUnauthorized,
		})
	}

	expTime := claims["expiresAt"].(int)

	// if expiration time is less than the
	// threshold time.
	if expTime < int(time.Now().Add(time.Minute*config.REFRESH_TOKEN_THRESHOLD).Unix()) {
		// get user id from valid parsed token
		userId := uint(int(claims["id"].(float64)))

		// refresh the JWT token
		refreshToken(userId, c)
	}

	// set user id parsed from JWT to the request object
	c.Locals("userId", fmt.Sprint(claims["id"]))

	return c.Next()
}

func refreshToken(id uint, c *fiber.Ctx) error {
	refreshToken, tokenizationErr := libs.GenerateToken(id)
	if tokenizationErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&outputs.StatusOutput{
			Status: fiber.StatusInternalServerError,
		})
	}

	// Renew token inside the cookie automatically
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * config.TOKEN_EXPIRATION_TIME),
		HTTPOnly: true,
	})

	return nil
}
