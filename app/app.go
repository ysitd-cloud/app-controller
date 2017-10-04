package app

func NewApplication(id string, environment Environment, information MetaInformation, scale AutoScale, network Network) Application {
	return &app{
		id:        id,
		info:      information,
		env:       environment,
		autoScale: scale,
		network:   network,
	}
}

func (app *app) GetID() string {
	return app.id
}

func (app *app) GetMetaInformation() MetaInformation {
	return app.info
}

func (app *app) GetEnvironment() Environment {
	return app.env
}

func (app *app) GetMeta() MetaInformation {
	return app.info
}

func (app *app) GetAutoScale() AutoScale {
	return app.autoScale
}

func (app *app) GetNetwork() Network {
	return app.network
}
