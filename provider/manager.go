package provider

import (
	"database/sql"

	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/app-controller/manager"
)

type managerServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*managerServiceProvider) Provides() []string {
	return []string{
		"manager",
	}
}

func (*managerServiceProvider) Register(app container.Container) {
	app.Bind("manager", func(app container.Container) interface{} {
		m := new(manager.Manager)
		db := app.Make("pg").(*sql.DB)
		m.SetDB(db)
		return m
	})
}
