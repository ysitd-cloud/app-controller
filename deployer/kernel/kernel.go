package kernel

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/deployer/provider"
	core "github.com/ysitd-cloud/app-controller/providers"
)

func CreateKernel() container.Kernel {
	kernel := container.NewKernel()
	kernel.Register(core.NewAzureServiceProvider)
	kernel.Register(core.NewDatabaseServiceProvider)
	kernel.Register(core.NewK8sServiceProvider)
	kernel.Register(core.NewEnvironmentManagerServiceProvider)
	kernel.Register(provider.NewClientServiceProvider)
	kernel.Register(provider.NewDeployerServiceProvider)

	return kernel
}
