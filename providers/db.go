package providers

import (
	"os"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tonyhhyip/go-di-container"
)

type databaseServiceProvider struct {
	*container.AbstractServiceProvider
}

func NewDatabaseServiceProvider(kernel container.Container) container.ServiceProvider {
	sp := databaseServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(kernel)

	return &sp
}

func (sp *databaseServiceProvider) Provides() []string {
	return []string{
		"db",
		"db.postgres",
		"db.postgres.url",
	}
}

func (sp *databaseServiceProvider) Register(kernel container.Container) {
	kernel.Instance("db.postgres.url", os.Getenv("DB_URL"))
	kernel.Bind("db.postgres", func(app container.Container) interface{} {
		dbUrl := kernel.Make("db.postgres.url").(string)
		db, err := sql.Open("postgres", dbUrl)
		if err != nil {
			panic(err)
		}

		return db
	})

	kernel.Alias("db", "db.postgres")
}
