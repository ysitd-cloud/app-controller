package app

type Application interface {
	GetID() string
	GetEnvironment() Environment
	GetMeta() MetaInformation
	GetAutoScale() AutoScale
	GetNetwork() Network
}

func NewApplication(id string, environment Environment, information MetaInformation, scale AutoScale, network Network) Application {
	return &app{
		id:        id,
		info:      information,
		env:       environment,
		autoScale: scale,
		network:   network,
	}
}

type app struct {
	id        string
	info      MetaInformation
	env       Environment
	autoScale AutoScale
	network   Network
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
