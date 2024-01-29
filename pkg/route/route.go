package route

import (
	"drr-server/app/controler"

	"github.com/gofiber/fiber/v2"
)

func Routes(route fiber.Router) {
	route.Get("/test", controler.HelloWorldGetCon)
	route.Post("/test", controler.HelloWorldPostCon)
}
