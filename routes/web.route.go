package routes

import (
	"bookstore/controllers"

	"github.com/gofiber/fiber/v2"
)

func WebRoutes(server *fiber.App) {
	server.Get("/", controllers.SendIndexWeb)
}
