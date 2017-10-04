package providers

import (
	"github.com/Azure/azure-storage-go"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/core"
)

type environmentManagerServiceProvider struct {
	*container.AbstractServiceProvider
}

func NewEnvironmentManagerServiceProvider(app container.Container) container.ServiceProvider {
	sp := environmentManagerServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)

	return &sp
}

func (sp *environmentManagerServiceProvider) Provides() []string {
	return []string{
		"core.manager.env",
	}
}

func (sp *environmentManagerServiceProvider) Register(app container.Container) {
	app.Bind("core.manager.env", func(app container.Container) interface{} {
		table := app.Make("azure.storage.table.app").(storage.Table)
		manager := core.NewEnvironmentManager(table)
		return manager
	})
}
