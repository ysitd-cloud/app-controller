package core

import (
	"github.com/ysitd-cloud/app-controller/app"
	"github.com/ysitd-cloud/app-controller/connect"
)

func NewEnvironmentManager() (EnvironmentManager, error) {
	table, err := connect.NewAzureTable()

	if err != nil {
		return nil, err
	}

	return &environmentManager{
		client: table,
	}, nil
}

func (e *environmentManager) GetEntry(id string) app.Environment {
	entity := e.client.GetEntityReference("env", id)

	env := make(app.Environment)

	for k, v := range entity.Properties {
		env.AddPair(k, v.(string))
	}

	return env
}
