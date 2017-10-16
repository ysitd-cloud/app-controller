package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMeta(c *gin.Context) {
	app := getApplication(c)
	meta := app.GetMeta()

	c.JSON(http.StatusOK, gin.H{
		"image": meta.GetImage(),
		"tag":   meta.GetTag(),
	})
}
