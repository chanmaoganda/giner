package routes

import (
	"log"
	"net/http"

	"github.com/chanmaoganda/giner/middleware"
	"github.com/chanmaoganda/giner/models"
	"github.com/gin-gonic/gin"
)

func LoginService(ctx *gin.Context) {
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
}