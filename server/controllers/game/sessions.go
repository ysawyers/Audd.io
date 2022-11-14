package game

import "github.com/gofiber/fiber/v2"

func FetchSessionState(c *fiber.Ctx) error {
	return c.JSON("fetch")
}

func CreateSession(c *fiber.Ctx) error {
	return c.JSON("create")
}

func DropSession(c *fiber.Ctx) error {
	// terminate in mongodb as well
	return c.JSON("delete")
}
