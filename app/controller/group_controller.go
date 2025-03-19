package controller

import (
	"github.com/ARTEM-kaa/ApiGO.git/app/model"
	"github.com/gofiber/fiber/v2"
)

func PostGroup(c *fiber.Ctx) error {
	return c.JSON(model.Group{})
}

func GetGroup(c *fiber.Ctx) error {
	return c.JSON(model.Group{})
}

func PutGroup(c *fiber.Ctx) error {
	return c.JSON(model.Group{})

}

func DeleteGroup(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
