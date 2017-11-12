package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNetwork(c *gin.Context) {
	app := getApplication(c)
	network := app.GetNetwork()

	c.JSON(http.StatusOK, gin.H{
		"domain": network.GetDomain(),
		"port":   network.GetPort(),
	})
}
