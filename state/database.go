package state

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var DB *sql.DB

func CreateDatabase() {
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

	DB = sql.OpenDB(database)

	err = DB.Ping()
	if err != nil {
		Log.Fatal("Database ping failed:", err)
	} else {
		Log.Debug("Database ping received!")
	}
}