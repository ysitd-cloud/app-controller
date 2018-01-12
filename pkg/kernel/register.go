package kernel

import (
	"code.ysitd.cloud/component/deployer/pkg/deployer"
	"code.ysitd.cloud/component/deployer/pkg/http"
	"code.ysitd.cloud/component/deployer/pkg/manager"
	"code.ysitd.cloud/component/deployer/pkg/service"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/go-common/db"
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
	Kernel.Register(db.NewServiceProvider)
	Kernel.Register(service.NewServiceProvider)
	Kernel.Register(http.NewServiceProvider)
}
