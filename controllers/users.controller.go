package controllers

import "github.com/gofiber/fiber/v2"

func SendHello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
