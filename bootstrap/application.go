package bootstrap

import (
	"github.com/chanmaoganda/giner/router"
	"github.com/gin-gonic/gin"
)

func Application() error {
	InitializeLogger()
	Log.Info("Logger initialized!")

	database := CreateDatabase()

	engine := gin.New()

	engine = router.MakeService(database, engine)

	address := "localhost:8080"

	return engine.Run(address)
}