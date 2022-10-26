package controller

import "github.com/gofiber/fiber/v2"

func GetAllCustomer(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"name":    "customer1",
		"surname": "customer1",
	})
}
