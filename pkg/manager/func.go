package manager

import "github.com/satori/go.uuid"

func genAppID() string {
	return uuid.NewV4().String()
}

func normalizeApplication(app *Application) *Application {
	if app.ID == "" {
		app.ID = genAppID()
	}

	return app
}
