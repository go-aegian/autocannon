package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sainzg/autocannon/webservers/go"
)

func setupRouter() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(weather.Predict(5))
	})

	return app
}

func main() {
	app := setupRouter()
	log.Fatal(app.Listen(":3007"))
}
