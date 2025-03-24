package service

import (
	"net/http"

	"github.com/chanmaoganda/giner/model"
	"github.com/chanmaoganda/giner/state"
	"github.com/gin-gonic/gin"
)

func FindByName(c *gin.Context) {
	// TODO: add redis logics here
	name := c.Param("username")
	user := model.QueryByName(state.DB, name)
	c.JSON(http.StatusOK, user)
}
