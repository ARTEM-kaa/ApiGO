package main

import (
	"log"

	"github.com/ARTEM-kaa/ApiGO.git/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	routes.ContactRoutes(app)
	routes.GroupRoutes(app)

	if err := app.Listen(":6080"); err != nil {
		log.Fatalln(err)
	}
}
