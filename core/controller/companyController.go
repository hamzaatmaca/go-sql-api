package controller

import "github.com/gofiber/fiber/v2"

func GetAllCompany(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{

		"companyId":   "",
		"companyName": "company1",
	})
}
