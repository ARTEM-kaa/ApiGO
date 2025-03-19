package controller

import (
	"github.com/ARTEM-kaa/ApiGO.git/app/model"
	"github.com/gofiber/fiber/v2"
)

func PostContact(c *fiber.Ctx) error {
	return c.JSON(model.Contact{})
}

func GetContact(c *fiber.Ctx) error {
	return c.JSON(model.Contact{})
}

func PutContact(c *fiber.Ctx) error {
	return c.JSON(model.Contact{})

}

func DeleteContact(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
