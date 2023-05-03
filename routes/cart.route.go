package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(api *fiber.Group) {
	user := api.Group("/cart")

	user.Get("/current", middleware.Protected, controllers.GetCurrentUserCart)
	user.Post("/add", middleware.Protected, controllers.AddToCart)
}
