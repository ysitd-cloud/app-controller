package kernel

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/tonyhhyip/go-di-container"
)

type postgresServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*postgresServiceProvider) Provides() []string {
	return []string{
		"pg",
		"pg.url",
	}
}

func (*postgresServiceProvider) Register(app container.Container) {
	app.Instance("pg.url", os.Getenv("DB_URL"))
	app.Bind("pg", func(app container.Container) interface{} {
		url := app.Make("pg.url").(string)
		db, err := sql.Open("postgres", url)
		if err != nil {
			panic(err)
		}

		return db
	})
}
