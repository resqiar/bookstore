package controllers

import "github.com/gofiber/fiber/v2"

func SendIndexWeb(c *fiber.Ctx) error {
	PATH := "./static/index/index.html"
	return c.SendFile(PATH)
}

func SendLoginWeb(c *fiber.Ctx) error {
	PATH := "./static/login/login.html"
	return c.SendFile(PATH)
}
