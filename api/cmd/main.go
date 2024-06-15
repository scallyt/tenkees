package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/test", func(c *fiber.Ctx) error {
		return c.JSON("ALma")
	})

	log.Fatal(app.Listen(":3000"))
}
