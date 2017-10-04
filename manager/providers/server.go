package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
)

func HttpServerServiceProviderBuilder(app container.Container) container.ServiceProvider {
	sp := HttpServerServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)

	return &sp
}

type HttpServerServiceProvider struct {
	*container.AbstractServiceProvider
}

func (sp *HttpServerServiceProvider) Register(app container.Container) {
	app.Singleton("http.server", func(app container.Container) interface{} {
		return gin.Default()
	})
}

func (sp *HttpServerServiceProvider) Provides() []string {
	return []string{"http.server"}
}
