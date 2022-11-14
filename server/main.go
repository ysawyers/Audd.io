package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
	"github.com/ysawyers/Audd.io/tree/master/server/routes"
)

func main() {
	fastergoding.Run()
	app := fiber.New()

	routes.PublicRoutes(app)

	app.Listen(":5000")
}
