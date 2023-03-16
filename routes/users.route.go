package routes

import (
	"bookstore/controllers"
	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(v1 *fiber.Group) {
	v1.Get("/", controllers.SendHello)
}
