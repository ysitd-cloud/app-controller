package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEnv(c *gin.Context) {
	app := getApplication(c)
	env := app.GetEnvironment()

	c.JSON(http.StatusOK, env)
}
