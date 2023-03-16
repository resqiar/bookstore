package middleware

import (
	"bookstore/libs"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// token extracted from bearer
	var token string

	if strings.HasPrefix(authHeader, "Bearer ") {
		// trim auth header to only get the value after "Bearer "
		token = strings.TrimPrefix(authHeader, "Bearer ")
	}

	// if there is no token inside the Authorization header,
	// send 401 status of unauthorized
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, valid := libs.ParseToken(token)
	if !valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// set user id parsed from JWT to the request object
	c.Locals("userId", fmt.Sprint(claims["id"]))

	return c.Next()
}
