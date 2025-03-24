package router

import (
	"github.com/chanmaoganda/giner/service"
	"github.com/gin-gonic/gin"
)

func MakeService() (*gin.Engine) {
	engine := gin.New()

	gin.SetMode(gin.ReleaseMode)
	
	v0 := engine.Group("/v0")

	v0.GET("/user/:username", service.FindByName)

	return engine
}
