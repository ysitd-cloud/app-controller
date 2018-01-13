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
	}
}

func (*managerServiceProvider) Register(app container.Container) {
	app.Bind("manager", func(app container.Container) interface{} {
		m := new(manager)
		db := app.Make("db.pool").(db.Pool)
		m.SetDB(db)
		return m
	})
}
