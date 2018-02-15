package http

import (
	"code.ysitd.cloud/grpc/schema/deployer"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
)

type httpServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*httpServiceProvider) Provides() []string {
	return []string{
		"http.server",
	}
}

func (*httpServiceProvider) Register(app container.Container) {
	app.Singleton("http.server", func(app container.Container) interface{} {
		service := app.Make("service").(deployer.DeployerServer)
		server := gin.Default()
		server.Use(func(c *gin.Context) {
			c.Set("service", service)
			c.Next()
		})
		register(server)
		return server
	})
}
