package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(api *fiber.Group) {
	api.Post("/book/add", middleware.Protected, middleware.Admin, controllers.AddBook)
	api.Get("/book/list", middleware.Protected, middleware.Admin, controllers.GetBooks)
}
