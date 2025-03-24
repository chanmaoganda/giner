package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Server(db *sql.DB) (*gin.Engine) {

	router := gin.Default()

	return router
}