package core

import (
	"github.com/Azure/azure-storage-go"
	"github.com/ysitd-cloud/app-controller/app"
)

func NewEnvironmentManager(table storage.Table) EnvironmentManager {
	return &environmentManager{
		client: table,
	}
}

func (e *environmentManager) GetEntry(id string) app.Environment {
	entity := e.client.GetEntityReference("env", id)

	env := make(app.Environment)

	for k, v := range entity.Properties {
		env.AddPair(k, v.(string))
	}

	return env
}
