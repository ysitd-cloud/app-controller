package core

import "github.com/ysitd-cloud/app-controller/app"

func (e *environmentManager) GetEntry(id string) app.Environment {
	entity := e.client.GetEntityReference("env", id)

	env := make(app.Environment)

	for k, v := range entity.Properties {
		env[k] = v.(string)
	}

	return env
}
