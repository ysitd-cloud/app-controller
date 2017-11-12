package provider

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/deployer"
	"k8s.io/client-go/kubernetes"
)

type deployerServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*deployerServiceProvider) Provides() []string {
	return []string{
		"deployer",
	}
}

func (*deployerServiceProvider) Register(app container.Container) {
	app.Bind("deployer", func(app container.Container) interface{} {
		c := new(deployer.Controller)
		client := app.Make("k8s.client").(*kubernetes.Clientset)
		c.SetClient(client)
		namespace := app.Make("k8s.namespace").(string)
		c.SetNamespace(namespace)
		return c
	})
}
