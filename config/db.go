package config

import (
	"Project-Workey/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dbURL := "postgresql://postgres:secret@:5432/postgres" //driver://postgres:password@localhost:5432/dbname
	dbURL := os.Getenv("DB_SOURCE")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("\nConnected to database:")

	db.AutoMigrate(&model.Login{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Worker{})
	db.AutoMigrate(&model.Verification{})
	return db
}
