package handler

import (
	"net/http"

	"code.ysitd.cloud/component/deployer/pkg/manager"
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/grpc-schema/deployer"
	"github.com/ysitd-cloud/grpc-schema/deployer/actions"
	"golang.org/x/net/context"
)

func CreateApplication(c *gin.Context) {
	var app manager.Application
	if err := c.BindJSON(&app); err != nil {
		c.AbortWithError(http.StatusPreconditionFailed, err)
		return
	}

	req := &actions.CreateApplicationRequest{
		App: app.ToPb(),
	}

	service := c.MustGet("service").(deployer.DeployerServer)
	reply, err := service.CreateApplication(context.Background(), req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, reply)
}

func GetApplicationByUsername(c *gin.Context) {
	user := c.Param("user")

	req := &actions.ListApplicationsByUsernameRequest{
		Username: user,
	}

	service := c.MustGet("service").(deployer.DeployerServer)
	reply, err := service.ListApplicationsByUsername(context.Background(), req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	apps := make([]*manager.Application, 0)
	for _, a := range reply.GetApps() {
		apps = append(apps, manager.FromPbToApplication(a))
	}
	c.AbortWithStatusJSON(http.StatusOK, apps)
}

func GetApplicationById(c *gin.Context) {
	id := c.Param("app")

	req := &actions.GetApplicationByIdRequest{
		Id: id,
	}

	service := c.MustGet("service").(deployer.DeployerServer)
	reply, err := service.GetApplicationById(context.Background(), req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if reply.GetExists() {
		c.AbortWithStatusJSON(http.StatusOK, manager.FromPbToApplication(reply.GetApp()))
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
