package main

import (
	"log"
	"regexp"

	"ejemplo.com/fiber/routes"
	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json" // libreria para hacer mas rapido el manejo de json
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// Texto validos para nombres y descripciones
var nombreRegex = regexp.MustCompile(`^[a-zA-Z0-9ñÑáéíóúÁÉÍÓÚ\s]+$`)

// Función de validación personalizada
func nombreValido(fl validator.FieldLevel) bool {
	return nombreRegex.MatchString(fl.Field().String())
}

func main() {

	validate := validator.New()
	validate.RegisterValidation("texto", nombreValido)

	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validate},
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
	})

	api := app.Group("/api")

	app.Get("/", handler)
	routes.TareasRoute(api)

	log.Fatal(app.Listen(":3000"))
}

func handler(c fiber.Ctx) error {
	return c.SendString("API Tareas")
}
