package routes

import (
	"ejemplo.com/fiber/handlers"
	"github.com/gofiber/fiber/v3"
)

func TareasRoute(route fiber.Router) {

	route.Get("/tareas", handlers.ListadoHandler)
	route.Get("/tareas/:tareaId", handlers.BuscarHandler)
	route.Post("/tareas", handlers.AlmacenarHandler)
	route.Put("/tareas", handlers.ActualizarHandler)
	route.Delete("/tareas", handlers.BorrarHandler)
}
