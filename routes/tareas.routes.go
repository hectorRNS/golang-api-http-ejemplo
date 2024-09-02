package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
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
	tareas := []*Tareas{
		{
			Tarea: "Tarea 1",
			Id:    1,
		},
		{
			Tarea: "Tarea 2",
			Id:    2,
		},
		{
			Tarea: "Tarea 33",
			Id:    3,
		},
	}

	return c.JSON(tareas)
}

func BuscarHandler(c fiber.Ctx) error {

	parametro := c.Params("tareaId")
	tareaId, err := strconv.Atoi(parametro) // convertir a entero

	if err != nil {
		return c.SendString("Parametro invalido")
	}

	var tareaEncontrada Tareas

	tareas := []*Tareas{
		{
			Tarea: "Tarea 1",
			Id:    1,
		},
		{
			Tarea: "Tarea 2",
			Id:    2,
		},
		{
			Tarea: "Tarea 33",
			Id:    3,
		},
	}

	for _, tarea := range tareas {

		if tarea.Id == tareaId {
			tareaEncontrada = *tarea
		}
	}

	return c.JSON(tareaEncontrada)
}

func handler(c fiber.Ctx) error {
	return c.SendString("Hello, World")
}
