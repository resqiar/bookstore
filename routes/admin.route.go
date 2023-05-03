package routes

import (
	"bookstore/controllers"
	"bookstore/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(api *fiber.Group) {
	api.Post("/adm/book/add", middleware.Protected, middleware.Admin, controllers.AddBook)
	api.Get("/adm/book/list", middleware.Protected, middleware.Admin, controllers.GetBooks)
	api.Post("/adm/book/edit/:id", middleware.Protected, middleware.Admin, controllers.GetBook)
	api.Post("/adm/book/delete/:id", middleware.Protected, middleware.Admin, controllers.DeleteBook)
}
