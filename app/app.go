package app

type Application interface {
	GetID() string
	GetEnvironment() Environment
	GetMeta() MetaInformation
}

type App struct {
	id       string
	info     MetaInformation
	env      Environment
	replicas uint
}

func (app *App) GetID() string {
	return app.id
}

func (app *App) GetMetaInformation() MetaInformation {
	return app.info
}

func (app *App) GetEnvironment() Environment {
	return app.env
}

func (app *App) GetReplicas() uint {
	return app.replicas
}
