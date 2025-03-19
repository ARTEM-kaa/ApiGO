package routes

import (
	"github.com/ARTEM-kaa/ApiGO.git/app/controller"
	"github.com/gofiber/fiber/v2"
)

func ContactRoutes(a *fiber.App) {
	route := a.Group("/api/v1/contact")

	route.Get("/", controller.GetContact)
	route.Post("/", controller.PostContact)
	route.Put("/", controller.PutContact)
	route.Delete("/", controller.DeleteContact)
}
