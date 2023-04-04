package controllers

import "github.com/gofiber/fiber/v2"

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
