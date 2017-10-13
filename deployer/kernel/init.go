package kernel

import (
	core "github.com/ysitd-cloud/app-controller/providers"
)

func init() {
	Kernel.Register(core.NewAzureServiceProvider)
	Kernel.Register(core.NewDatabaseServiceProvider)
	Kernel.Register(core.NewK8sServiceProvider)
	Kernel.Register(core.NewEnvironmentManagerServiceProvider)
}
