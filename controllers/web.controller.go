package controllers

import "github.com/gofiber/fiber/v2"

func SendIndexWeb(c *fiber.Ctx) error {
	PATH := "./static/index/index.html"
	return c.SendFile(PATH)
}
