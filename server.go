package main

import (
	"log"

	"github.com/go-sql-api/core/router"
	"github.com/go-sql-api/database"
	"github.com/go-sql-api/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.EstablishedDBConnect()
	app := fiber.New()
	app.Use(cors.New())

	router.AllRoutes(app)

	log.Fatal(app.Listen(":" + env.ENV("PORT")))

}
