package providers

import (
	"os"
	"path/filepath"

	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/gin-utils/env"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sServiceProvider struct {
	*container.AbstractServiceProvider
}

func NewK8sServiceProvider(app container.Container) container.ServiceProvider {
	sp := k8sServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)

	return &sp
}

func (sp *k8sServiceProvider) Provides() []string {
	return []string{
		"k8s.config",
		"k8s.client",
		"k8s.namespace",
	}
}

func (sp *k8sServiceProvider) Register(app container.Container) {
	app.Instance("k8s.namespace", env.GetEnvWithDefault("K8S_NAMESPACE", "app"))
	app.Singleton("k8s.config", func(app container.Container) interface{} {
		config, err := rest.InClusterConfig()
		if err == nil {
			return config
		}

		file := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", file)
		if err != nil {
			panic(err)
		}

		return config
	})

	app.Bind("k8s.client", func(app container.Container) interface{} {
		config := app.Make("k8s.config").(*rest.Config)
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err)
		}

		return client
	})
}
