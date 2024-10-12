package main

import (
	"log"

	"ejemplo.com/fiber/routes"
	"ejemplo.com/fiber/validaciones"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json" // libreria para hacer mas rapido el manejo de json
	"github.com/gofiber/fiber/v3"
)

type StructValidator struct {
	validate *validator.Validate
}

func (v *StructValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	validate := validator.New()
	validaciones.RegistrarValiaciones(validate)

	app := fiber.New(fiber.Config{
		StructValidator: &StructValidator{validate: validate},
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
	})

	api := app.Group("/api")

	app.Get("/", handler)
	routes.TareasRoute(api)

	// HTTP
	log.Fatal(app.Listen(":3000"))

	// * para agregar TLS, HTTPS
	// log.Fatal(app.Listen(":3001", fiber.ListenConfig{
	// 	CertFile:    "certs/server.crt",
	// 	CertKeyFile: "certs/server.key",
	// }))
}

func handler(c fiber.Ctx) error {
	return c.SendString("API Tareas")
}
