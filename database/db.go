package database

import (
	"log"
	"os"

	"API/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Println(err)
		return nil, err
	}
	DB = db
	return DB, nil
}
