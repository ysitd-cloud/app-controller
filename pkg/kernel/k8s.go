package kernel

import (
	"os"

	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/k8s-utils"
)

type k8sServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*k8sServiceProvider) Provides() []string {
	return []string{
		"k8s.namespace",
		"k8s.client",
	}
}

func (*k8sServiceProvider) Register(app container.Container) {
	app.Instance("k8s.namespace", os.Getenv("APP_NAMESPACE"))
	app.Bind("k8s.client", func(app container.Container) interface{} {
		client, err := utils.AutoConnect()
		if err != nil {
			panic(err)
		}

		return client
	})
}
