package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TareasRoute(route fiber.Router) {

	route.Get("/tareas", ListadoHandler)
	route.Get("/tareas/:tareaId", BuscarHandler)
	route.Post("/tareas", handler)
	route.Put("/tareas", handler)
	route.Delete("/tareas", handler)
}

type Tareas struct {
	Id    int    `json:"id"`
	Tarea string `json:"tarea"`
}

func ListadoHandler(c fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	var tareas []Tareas
	db.Limit(6).Find(&tareas)

	return c.JSON(tareas)
}

func BuscarHandler(c fiber.Ctx) error {

	parametro := c.Params("tareaId")
	tareaId, err := strconv.Atoi(parametro) // convertir a entero

	if err != nil {
		return c.SendString("Parametro invalido")
	}

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	var tareas Tareas
	db.First(&tareas, tareaId)

	return c.JSON(tareas)
}

func handler(c fiber.Ctx) error {
	return c.SendString("Hello, World")
}
