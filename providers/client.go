package providers

import (
	"github.com/tonyhhyip/go-di-container"
	"k8s.io/client-go/kubernetes"
)

type clientServiceProvider struct {
	*container.AbstractServiceProvider
}

func NewClientServiceProvider(kernel container.Container) container.ServiceProvider {
	sp := clientServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(kernel)

	return &sp
}

func (*clientServiceProvider) Provides() []string {
	return []string{
		"k8s.client.deployment",
		"k8s.client.ingress",
		"k8s.client.service",
		"k8s.client.secret",
	}
}

func (*clientServiceProvider) Register(app container.Container) {
	app.Bind("k8s.client.deployment", func(app container.Container) interface{} {
		client := app.Make("k8s.client").(*kubernetes.Clientset)
		namespace := app.Make("k8s.namespace").(string)
		deploymentClient := client.AppsV1beta1().Deployments(namespace)
		return deploymentClient
	})
	app.Bind("k8s.client.ingress", func(app container.Container) interface{} {
		namespace := app.Make("k8s.namespace").(string)
		client := app.Make("k8s.client").(*kubernetes.Clientset)
		ingressClient := client.ExtensionsV1beta1().Ingresses(namespace)
		return ingressClient
	})
	app.Bind("k8s.client.service", func(app container.Container) interface{} {
		client := app.Make("k8s.client").(*kubernetes.Clientset)
		namespace := app.Make("k8s.namespace").(string)
		serviceClient := client.CoreV1().Services(namespace)
		return serviceClient
	})
	app.Bind("k8s.client.secret", func(app container.Container) interface{} {
		client := app.Make("k8s.client").(*kubernetes.Clientset)
		namespace := app.Make("k8s.namespace").(string)
		secretClient := client.CoreV1().Secrets(namespace)
		return secretClient
	})
}
