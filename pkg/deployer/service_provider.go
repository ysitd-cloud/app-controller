package deployer

import (
	"github.com/tonyhhyip/go-di-container"
	"k8s.io/client-go/kubernetes"
)

type deployerServiceProvider struct {
	*container.AbstractServiceProvider
}

func CreateDeployerServiceProvider(app container.Container) container.ServiceProvider {
	sp := &deployerServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)
	return sp
}

func (*deployerServiceProvider) Provides() []string {
	return []string{
		"deployer",
	}
}

func (*deployerServiceProvider) Register(app container.Container) {
	app.Bind("deployer", func(app container.Container) interface{} {
		c := new(controller)
		client := app.Make("k8s.client").(kubernetes.Interface)
		c.SetClient(client)
		namespace := app.Make("k8s.namespace").(string)
		c.SetNamespace(namespace)
		return c
	})
}
