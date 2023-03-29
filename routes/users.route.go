package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(api *fiber.Group) {
	user := api.Group("/user")

	user.Get("/current", middleware.Protected, controllers.GetCurrentUser)
}
