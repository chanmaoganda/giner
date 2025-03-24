package bootstrap

import (
	"github.com/chanmaoganda/giner/router"
	"github.com/chanmaoganda/giner/state"
	"github.com/gin-gonic/gin"
)

func Application() error {
	gin.SetMode(gin.ReleaseMode)

	state.InitializeLogger()

	state.CreateDatabase()

	engine := router.MakeService()

	address := "localhost:8080"

	state.Log.Info("Currently running on " + address)

	return engine.Run(address)
}