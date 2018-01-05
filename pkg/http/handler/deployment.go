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

	id := c.Param("app")
	var deployment manager.Deployment

	if c.BindJSON(&deployment) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	confirm, e, err := m.UpdateDeployment(id, &deployment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ok := false

	defer func() {
		confirm <- ok
		if ok {
			err := <-e
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			} else {
				c.AbortWithStatus(http.StatusOK)
			}
		}
		close(confirm)
	}()

	d := kernel.Make("deployer").(deployer.Controller)
	if _, err := d.UpdateDeploymentImage(template.GetName(id), deployment.Image, deployment.Tag); err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ok = true
}
