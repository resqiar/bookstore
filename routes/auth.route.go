package routes

import (
	"bookstore/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api *fiber.Group) {
	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/logout", controllers.Logout)
}
