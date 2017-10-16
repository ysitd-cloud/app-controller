package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/app"
	"github.com/ysitd-cloud/app-controller/core"
)

func getApplication(c *gin.Context) app.Application {
	kernel := c.MustGet("kernel").(container.Kernel)
	manager := kernel.Make("manager").(core.Manager)
	defer manager.Close()

	id := c.Param("id")
	app := manager.GetApplication(id)
	return app
}
