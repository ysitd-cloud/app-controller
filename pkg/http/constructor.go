package http

import "github.com/tonyhhyip/go-di-container"

func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &httpServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)

	return sp
}
