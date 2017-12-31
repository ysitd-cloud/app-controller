package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/pkg/deployer"
	"github.com/ysitd-cloud/app-controller/pkg/manager"
	"github.com/ysitd-cloud/app-controller/pkg/template"
)

func UpdateImage(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	m := kernel.Make("manager").(manager.Manager)
	defer m.Close()

	id := c.Param("app")
	var deployment manager.Deployment

	if c.BindJSON(&deployment) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := m.UpdateDeployment(id, &deployment); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	d := kernel.Make("deployer").(deployer.Controller)
	if _, err := d.UpdateDeploymentImage(template.GetName(id), deployment.Image, deployment.Tag); err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
