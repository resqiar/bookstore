package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api *fiber.Group) {
	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	// Test protected route, should not be accessed
	// without providing credential.
	auth.Get("/protected", middleware.Protected, func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)
		return c.SendString("Your user id = " + userId)
	})
}
