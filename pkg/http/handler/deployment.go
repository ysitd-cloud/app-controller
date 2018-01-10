package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/grpc-schema/deployer"
	"github.com/ysitd-cloud/grpc-schema/deployer/actions"
	"github.com/ysitd-cloud/grpc-schema/deployer/models"
	"golang.org/x/net/context"
)

func UpdateImage(c *gin.Context) {
	id := c.Param("app")
	var deployment models.Deployment

	if c.BindJSON(&deployment) != nil {
		c.AbortWithStatus(http.StatusPreconditionFailed)
		return
	}

	req := &actions.UpdateDeploymentImageRequest{
		Id:         id,
		Deployment: &deployment,
	}

	service := c.MustGet("service").(deployer.DeployerServer)
	reply, err := service.UpdateDeploymentImage(context.Background(), req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.AbortWithStatusJSON(http.StatusCreated, reply)
}
