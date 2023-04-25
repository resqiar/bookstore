package routes

import (
	"bookstore/controllers"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(api *fiber.Group) {
	api.Get("/book/list", controllers.GetBooks)
}
