package manager

import (
	"database/sql"

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
		m := new(Manager)
		db := app.Make("pg").(*sql.DB)
		m.SetDB(db)
		return m
	})
}
