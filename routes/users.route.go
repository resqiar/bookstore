package routes

import (
	"bookstore/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(api *fiber.Group) {
	api.Get("/", controllers.SendHello)
}
