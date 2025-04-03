package giner

import (
	"log"
	"net/http"

	"github.com/chanmaoganda/giner/middleware"
	"github.com/chanmaoganda/giner/models"
	"github.com/gin-gonic/gin"
)

func PrepareLogs() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[DEBUG] ")
}

func PrepareAppStates() {
	InitDataBase()
	
}

func Run() {
	PrepareLogs()

	LoadEnv()

	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		var user models.AuthUser
		if ctx.BindJSON(&user) == nil {
			token, err := middleware.SignToken(user.Username)
			if err != nil {
				log.Println("[DEBUG] ", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "Signing JWT Error",
				})
				return
			}
			log.Println("sign token successful")

			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		}
	})

	router.GET("/ping", middleware.AuthMiddleWareJWT(), func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
