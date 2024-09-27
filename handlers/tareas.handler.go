package handlers

import (
	"strconv"

	"ejemplo.com/fiber/structs"
	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ListadoHandler(c fiber.Ctx) error {

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	var tareas []structs.Tareas
	db.Limit(6).Find(&tareas)

	return c.JSON(tareas)
}

func BuscarHandler(c fiber.Ctx) error {

	parametro := c.Params("tareaId")
	tareaId, err := strconv.Atoi(parametro) // convertir a entero

	if err != nil {
		return c.SendString("Parametro invalido")
	}

	// log.Info(tareaId)

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	var tareas structs.Tareas
	db.First(&tareas, tareaId)

	return c.JSON(tareas)
}

func AlmacenarHandler(c fiber.Ctx) error {
	request := new(structs.NuevaTarea)

	if err := c.Bind().Body(request); err != nil {
		return err
	}

	// log.Println(request.Tarea)

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	tarea := structs.Tareas{Tarea: request.Tarea}
	db.Create(&tarea)

	return c.SendString("Creado correctamente")
}

func ActualizarHandler(c fiber.Ctx) error {

	request := new(structs.Tareas)

	if err := c.Bind().Body(request); err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	db.Save(&request).Limit(1)

	return c.SendString("Actualizado correctamente")
}

func BorrarHandler(c fiber.Ctx) error {
	request := new(structs.TareaId)

	if err := c.Bind().Body(request); err != nil {
		return err
	}

	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return c.SendString("error conexion base de datos")
	}

	db.Delete(&structs.Tareas{}, request.TareaId).Limit(1)

	return c.SendString("Borrado correctamente")
}
