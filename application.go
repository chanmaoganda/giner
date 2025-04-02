package giner

import (
	"log"
	"net/http"

	"github.com/chanmaoganda/giner/middleware"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
}

func Run() {
	InitDataBase()

	router := gin.Default()

	router.POST("/login", func(ctx *gin.Context) {
		var user User
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
