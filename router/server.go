package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func MakeService(db *sql.DB, engine *gin.Engine) (*gin.Engine) {

	return engine
}