package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func WebRoutes(server *fiber.App) {
	server.Get("/", controllers.SendIndexWeb)
	server.Get("/login", controllers.SendLoginWeb)
	server.Get("/register", controllers.SendRegisterWeb)

	server.Get("/browse", controllers.SendBrowseWeb)
	server.Get("/cart", middleware.Protected, controllers.SendCartWeb)

	// BOOK ROUTE
	server.Get("/book/:id", controllers.SendBookDetailWeb)

	// ADMIN ROUTES
	admin := server.Group("/admin", middleware.Protected, middleware.Admin)
	admin.Get("/", controllers.SendAdminIndex)
	admin.Get("/book/add", controllers.SendAdminAddBook)
	admin.Get("/book/list", controllers.SendAdminListBook)
	admin.Get("/book/edit/:id", controllers.SendAdminEditBook)
}
