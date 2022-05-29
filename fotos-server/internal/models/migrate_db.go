package models

import (
	"github.com/SuhailAhmed2627/fotos-server/database"
)

func MigrateDB() {
	db := database.GetDB()

	for _, model := range []interface{}{
		UserRegistration{},
		User{},
		Event{},
		Image{},
		Face{},
	} {
		db.AutoMigrate(&model)
	}
}
