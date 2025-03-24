package bootstrap

import (
	"github.com/chanmaoganda/giner/router"
)

func Application() error {
	database := CreateDatabase()

	router := router.Server(database)

	address := "localhost:8080"

	return router.Run(address)
}