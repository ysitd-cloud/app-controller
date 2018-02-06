package manager

import (
	"code.ysitd.cloud/common/go/db"
	"github.com/tonyhhyip/go-di-container"
)

type managerServiceProvider struct {
	*container.AbstractServiceProvider
}

func CreateManagerServiceProvider(app container.Container) container.ServiceProvider {
	sp := &managerServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)

	return sp
}

func (*managerServiceProvider) Provides() []string {
	return []string{
		"manager",
		"manager.app",
		"manager.deployment",
		"manager.env",
		"manager.network",
	}
}

func (*managerServiceProvider) Register(app container.Container) {
	app.Bind("manager.app", func(app container.Container) interface{} {
		db := app.Make("db.pool").(db.Pool)
		store := new(appStore)
		store.db = db
		return store
	})

	app.Bind("manager.deployment", func(app container.Container) interface{} {
		db := app.Make("db.pool").(db.Pool)
		store := new(deploymentStore)
		store.db = db
		return store
	})

	app.Bind("manager.env", func(app container.Container) interface{} {
		db := app.Make("db.pool").(db.Pool)
		store := new(envStore)
		store.db = db
		return store
	})

	app.Bind("manager.network", func(app container.Container) interface{} {
		db := app.Make("db.pool").(db.Pool)
		store := new(networkStore)
		store.db = db
		return store
	})

	app.Bind("manager", func(app container.Container) interface{} {
		as := app.Make("manager.app").(*appStore)
		ds := app.Make("manager.deployment").(*deploymentStore)
		es := app.Make("manager.env").(*envStore)
		ns := app.Make("manager.network").(*networkStore)
		m := &manager{
			app:        as,
			deployment: ds,
			env:        es,
			network:    ns,
		}
		return m
	})
}
