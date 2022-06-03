package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct{}

//DoSomething will do something
func (y controller) DoSomething(c *gin.Context) {
	c.JSON(http.StatusOK, true)
}
