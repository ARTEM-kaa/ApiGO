package routes

import (
	"github.com/ARTEM-kaa/ApiGO.git/app/controller"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(a *fiber.App) {
	routes := a.Group("/api/v1/group")

	routes.Get("/", controller.GetGroup)
	routes.Post("/", controller.PostGroup)
	routes.Put("/", controller.PutGroup)
	routes.Delete("/", controller.DeleteGroup)

}
