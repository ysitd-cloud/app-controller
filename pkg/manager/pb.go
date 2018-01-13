package manager

import "code.ysitd.cloud/grpc/schema/deployer/models"

func FromPbToApplication(app *models.Application) *Application {
	return &Application{
		ID:          app.GetId(),
		Owner:       app.GetOwner(),
		Name:        app.GetName(),
		Deployment:  app.GetDeployment(),
		Network:     app.GetNetwork(),
		Environment: FromPbToEnvironment(app.Environment),
	}
}

func FromPbToEnvironment(pairs []*models.EnvironmentPair) (env Environment) {
	env = make(Environment)
	for _, pair := range pairs {
		env[pair.GetKey()] = pair.GetValue()
	}
	return
}

func (app *Application) ToPb() *models.Application {
	return &models.Application{
		Id:          app.ID,
		Owner:       app.Owner,
		Name:        app.Name,
		Deployment:  app.Deployment,
		Network:     app.Network,
		Environment: app.Environment.ToPb(),
	}
}

func (env Environment) ToPb() (pairs []*models.EnvironmentPair) {
	pairs = make([]*models.EnvironmentPair, 0)
	for k, v := range env {
		pairs = append(pairs, &models.EnvironmentPair{
			Key:   k,
			Value: v,
		})
	}
	return
}
