package core

import "github.com/satori/go.uuid"

type App struct {
	manager *AppManager
	id uuid.UUID
	name string
	image string
	port uint32
	hostname string

	metaModified bool
	serviceModified bool
	deployModified bool
}

func (app *App) SetName(name string) {
	app.name = name
	app.metaModified = true
}

