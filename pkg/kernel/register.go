package kernel

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/pkg/deployer"
	"github.com/ysitd-cloud/app-controller/pkg/manager"
)

func init() {
	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &k8sServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return sp
	})

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &postgresServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return sp
	})

	Kernel.Register(deployer.CreateDeployerServiceProvider)
	Kernel.Register(manager.CreateManagerServiceProvider)
}
