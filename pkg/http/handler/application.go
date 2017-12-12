package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/pkg/deployer"
	"github.com/ysitd-cloud/app-controller/pkg/manager"
	"github.com/ysitd-cloud/app-controller/pkg/template"
)

func CreateApplication(c *gin.Context) {
	var app manager.Application
	if err := c.BindJSON(&app); err != nil {
		c.AbortWithError(http.StatusPreconditionFailed, err)
		return
	}

	kernel := c.MustGet("kernel").(container.Kernel)
	m := kernel.Make("manager").(manager.Manager)
	defer m.Close()

	if err := m.CreateApplication(app); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	d := kernel.Make("deployer").(deployer.Controller)

	secret := template.GenerateSecret(app.ID, app.Environment)
	if _, err := d.CreateSecret(secret); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	deployment := app.Deployment
	k8sDeployment := template.GenerateDeployment(app.ID, deployment.Image, deployment.Tag, app.Environment)
	if _, err := d.CreateDeployment(k8sDeployment); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	service := template.GenerateService(app.ID)
	if _, err := d.CreateService(service); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	network := app.Network
	ingress := template.GenerateIngress(app.ID, network.Domain)
	if _, err := d.CreateIngress(ingress); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}
