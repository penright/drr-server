package main

import (
	"drr-server/pkg/config"
	"drr-server/pkg/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.FiberConfig()
	app := fiber.New(config)
	app.Static("/", "./angular/browser")
	api := app.Group("/api")
	route.Routes(api)
	app.Get("/*", func(c *fiber.Ctx) error { return c.SendFile("./angular/browser/index.html") })
	app.Listen(":9091")
}
