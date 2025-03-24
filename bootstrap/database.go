package bootstrap

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func CreateDatabase() (*sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading [.env] file")
	}

	database_url := os.Getenv("DATABASE_URL")
	log.Println(database_url)

	database, err := pq.NewConnector(database_url)
	if err != nil {
		log.Fatal("Error connecting to " + database_url)
	}

	db := sql.OpenDB(database)

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	return db
}