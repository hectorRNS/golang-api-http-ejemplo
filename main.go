package main

import (
	"log"

	"ejemplo.com/fiber/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	app.Get("/", handler)
	routes.TareasRoute(api)

	log.Fatal(app.Listen(":3000"))
}

func handler(c fiber.Ctx) error {
	return c.SendString("API Tareas")
}
