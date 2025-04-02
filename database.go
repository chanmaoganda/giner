package giner

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // Import pq driver
)

var DB *sql.DB

func InitDataBase() {
	database_url := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", database_url)

	if err != nil {
		log.Panicln("Database open failed")
	}
	
	DB = db
}
