package giner

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	database_url := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(database_url), &gorm.Config{})

	if err != nil {
		log.Panicln("Database open failed")
	}
	
	DB = db
}
