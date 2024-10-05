package dbConexion

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Conectar() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("tareas.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
