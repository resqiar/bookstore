package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(v1 *fiber.Group) {
	auth := v1.Group("/auth")

	auth.Post("/login", controllers.Login)

	// Test protected route, should not be accessed
	// without providing credential.
	auth.Get("/protected", middleware.Protected, func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)
		return c.SendString("Your user id = " + userId)
	})
}
