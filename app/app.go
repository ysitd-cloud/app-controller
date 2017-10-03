package app

type Application interface {
	GetID() string
	GetEnvironment() Environment
	GetMeta() MetaInformation
	GetAutoScale() AutoScale
	GetNetwork() Network
}

func NewApplication(id string, environment Environment, information MetaInformation, scale AutoScale, network Network) Application {
	return &App{
		id:        id,
		info:      information,
		env:       environment,
		autoScale: scale,
		network:   network,
	}
}

type App struct {
	id        string
	info      MetaInformation
	env       Environment
	autoScale AutoScale
	network   Network
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

func (app *App) GetMeta() MetaInformation {
	return app.info
}

func (app *App) GetAutoScale() AutoScale {
	return app.autoScale
}

func (app *App) GetNetwork() Network {
	return app.network
}
