package controllers

import (
	"bookstore/database"
	"bookstore/entities"

	"github.com/gofiber/fiber/v2"
)

func SendIndexWeb(c *fiber.Ctx) error {
	PATH := "./static/dashboard/index.html"
	return c.SendFile(PATH)
}

func SendLoginWeb(c *fiber.Ctx) error {
	PATH := "./static/login/login.html"
	return c.SendFile(PATH)
}

func SendRegisterWeb(c *fiber.Ctx) error {
	PATH := "./static/register/register.html"
	return c.SendFile(PATH)
}

func SendBrowseWeb(c *fiber.Ctx) error {
	PATH := "./static/browse-book/browse.html"
	return c.SendFile(PATH)
}

func SendCartWeb(c *fiber.Ctx) error {
	PATH := "./static/cart/cart.html"
	return c.SendFile(PATH)
}

/**
 * BOOK ROUTES *
 **/
func SendBookDetailWeb(c *fiber.Ctx) error {
	id := c.Params("id")

	var book entities.Book

	result := database.DB.First(&book, "id = ?", id)
	if result.Error != nil {
		NOT_FOUND_PATH := "./static/404/404.html"
		return c.Status(fiber.StatusNotFound).SendFile(NOT_FOUND_PATH)
	}

	PATH := "./static/book/detail.html"
	return c.Status(fiber.StatusOK).SendFile(PATH)
}

/**
 * ADMIN ROUTES *
 **/
func SendAdminIndex(c *fiber.Ctx) error {
	PATH := "./static/admin/dashboard.html"
	return c.SendFile(PATH)
}

func SendAdminAddBook(c *fiber.Ctx) error {
	PATH := "./static/admin/book/add-book.html"
	return c.SendFile(PATH)
}

func SendAdminListBook(c *fiber.Ctx) error {
	PATH := "./static/admin/book/list-book.html"
	return c.SendFile(PATH)
}
