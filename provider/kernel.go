package provider

import "github.com/tonyhhyip/go-di-container"

var Kernel container.Kernel

func init() {
	Kernel = container.NewKernel()

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

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &deployerServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}

		sp.SetContainer(app)
		return sp
	})

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &managerServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return sp
	})
}
