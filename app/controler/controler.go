package controler

import (
	"drr-server/app/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloWorldGetCon(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}
func HelloWorldPostCon(c *fiber.Ctx) error {
	tmpHello := model.Hello{}
	err := c.BodyParser(&tmpHello)
	if err != nil {
		log.Println("Could not parse body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error":   "Could not parse body",
			"Message": err.Error(),
		})
	}
	log.Println(tmpHello.Name + ", " + tmpHello.Work)
	// do some work
	tmpNewHello := fmt.Sprintf("Hello %s is at %s",
		tmpHello.Name, tmpHello.Work)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"hello": tmpNewHello,
	})
}
