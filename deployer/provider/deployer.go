package provider

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/core"
	appv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func NewDeployerServiceProvider(app container.Container) container.ServiceProvider {
	sp := deployerServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)

	return &sp
}

type deployerServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*deployerServiceProvider) Provides() []string {
	return []string{
		"core.deployer",
	}
}

func (*deployerServiceProvider) Register(app container.Container) {
	app.Bind("core.deployer", func(app container.Container) interface{} {
		deployment := app.Make("k8s.client.deployment").(appv1beta1.DeploymentInterface)
		ingress := app.Make("k8s.client.ingress").(extv1beta1.IngressInterface)
		secret := app.Make("k8s.client.secret").(corev1.SecretInterface)
		service := app.Make("k8s.client.service").(corev1.ServiceInterface)
		manager := app.Make("core.manager").(core.Manager)

		deployer := core.NewKubernetesDeployer(manager, deployment, ingress, service, secret)
		return deployer
	})
}
