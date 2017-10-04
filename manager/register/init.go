package register

import (
	"github.com/ysitd-cloud/app-controller/manager/providers"
	core "github.com/ysitd-cloud/app-controller/providers"
)

func init() {
	Kernel.Register(providers.HttpServerServiceProviderBuilder)
	Kernel.Register(core.NewAzureServiceProvider)
	Kernel.Register(core.NewDatabaseServiceProvider)
	Kernel.Register(core.NewK8sServiceProvider)
	Kernel.Register(core.NewEnvironmentManagerServiceProvider)
}
