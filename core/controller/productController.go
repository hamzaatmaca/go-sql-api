package controller

import "github.com/gofiber/fiber/v2"

func GetAllProduct(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    "all products",
	})
}
