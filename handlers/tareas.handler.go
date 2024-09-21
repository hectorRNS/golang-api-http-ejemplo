package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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

type NuevaTarea struct {
	Tarea string `json:"tarea" xml:"tarea" form:"tarea"`
}

func AlmacenarHandler(c fiber.Ctx) error {
	request := new(NuevaTarea)

	if err := c.Bind().Body(request); err != nil {
		return err
	}

	// log.Println(request.Tarea)

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	tarea := Tareas{Tarea: request.Tarea}
	db.Create(&tarea)

	return c.SendString("Creado correctamente")
}

func Handler(c fiber.Ctx) error {
	return c.SendString("Hello, World")
}
