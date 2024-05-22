package database

import (
	"log"

	"github.com/sawalreverr/recything/internal/user/entity"
)

func AutoMigrate(db Database) {
	if err := db.GetDB().AutoMigrate(
		&entity.User{},
	); err != nil {
		log.Fatal("Database Migration Failed!")
	}

	log.Println("Database Migration Success")
}
