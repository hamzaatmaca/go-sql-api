package router

import (
	"github.com/go-sql-api/auth"
	"github.com/go-sql-api/core/controller"
	"github.com/go-sql-api/core/middleware"
	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App) {

	api := app.Group("/api/v1/")

	authenticate := api.Group("/login")
	authenticate.Post("/", auth.Login)

	productWay := api.Group("/product")
	productWay.Get("/", middleware.Protected(), controller.GetAllProduct)

	customerWay := api.Group("/customer")
	customerWay.Get("/", middleware.Protected(), controller.GetAllCustomer)

	companyWay := api.Group("/company")
	companyWay.Get("/", middleware.Protected(), controller.GetAllCompany)
}
