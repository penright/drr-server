package main

import (
	"drr-server/pkg/config"
	"drr-server/pkg/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.FiberConfig()
	app := fiber.New(config)
	api := app.Group("/api")
	route.Routes(api)
	app.Listen(":3000")
}
