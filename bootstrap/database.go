package bootstrap

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func CreateDatabase() (*sql.DB) {
	err := godotenv.Load()
	if err != nil {
		Log.Fatal("Error Loading [.env] file")
	}

	database_url := os.Getenv("DATABASE_URL")
	Log.Debug(database_url)

	database, err := pq.NewConnector(database_url)
	if err != nil {
		Log.Fatal("Error connecting to " + database_url)
	}

	db := sql.OpenDB(database)

	err = db.Ping()
	if err != nil {
		Log.Fatal("Database ping failed:", err)
	}

	return db
}