package service

import (
	"code.ysitd.cloud/component/deployer/pkg/deployer"
	"code.ysitd.cloud/component/deployer/pkg/manager"
	"code.ysitd.cloud/component/deployer/pkg/template"
	"code.ysitd.cloud/grpc/schema/deployer/actions"
	"code.ysitd.cloud/grpc/schema/deployer/models"
	"golang.org/x/net/context"
)

func (s *service) deploymentController() deployer.DeploymentController {
	return s.deployer.GetDeploymentController()
}

func (s *service) secretController() deployer.SecretController {
	return s.deployer.GetSecretController()
}

func (s *service) serviceController() deployer.ServiceController {
	return s.deployer.GetServiceController()
}

func (s *service) ingressController() deployer.IngressController {
	return s.deployer.GetIngressController()
}

func (s *service) ListApplicationsByUsername(_ context.Context, req *actions.ListApplicationsByUsernameRequest) (reply *actions.ListApplicationsByUsernameReply, err error) {
	username := req.GetUsername()
	store := s.manager.GetApplicationStore()
	apps, err := store.GetByOwner(username)
	if err != nil {
		return
	}

	pbApps := make([]*models.Application, 0)
	for _, app := range apps {
		pbApps = append(pbApps, app.ToPb())
	}
	reply = &actions.ListApplicationsByUsernameReply{
		Apps: pbApps,
	}
	return
}

func (s *service) CreateApplication(_ context.Context, req *actions.CreateApplicationRequest) (reply *actions.CreateApplicationReply, err error) {
	app := manager.FromPbToApplication(req.App)
	store := s.manager.GetApplicationStore()
	phase, err := store.Create(app)
	if err != nil {
		return
	}

	env := app.Environment
	secret := template.GenerateSecret(app.ID, env)
	if _, err = s.secretController().Create(secret); err != nil {
		phase.Cancel()
		return
	}

	d := app.Deployment
	deployment := template.GenerateDeployment(app.ID, d.Image, d.Tag, env)
	if _, err = s.deploymentController().Create(deployment); err != nil {
		phase.Cancel()
		return
	}

	service := template.GenerateService(app.ID)
	if _, err = s.serviceController().Create(service); err != nil {
		phase.Cancel()
		return
	}

	ingress := template.GenerateIngress(app.ID, app.Network.GetDomain())
	if _, err = s.ingressController().Create(ingress); err != nil {
		phase.Cancel()
		return
	}

	err = phase.Ok()
	if err != nil {
		return
	}

	reply = &actions.CreateApplicationReply{
		Success: true,
	}

	return
}

func (s *service) UpdateDeploymentImage(_ context.Context, req *actions.UpdateDeploymentImageRequest) (reply *actions.UpdateDeploymentImageReply, err error) {
	deployment := req.GetDeployment()
	id := req.GetId()

	phase, err := s.manager.GetDeploymentStore().Update(id, deployment)
	if err != nil {
		return
	}

	_, err = s.deploymentController().UpdateImage(template.GetName(id), deployment.GetImage(), deployment.GetTag())
	if err != nil {
		phase.Cancel()
		return
	}

	err = phase.Ok()
	if err != nil {
		return
	}

	reply = &actions.UpdateDeploymentImageReply{
		Success: true,
	}
	return
}

func (s *service) GetApplicationById(_ context.Context, req *actions.GetApplicationByIdRequest) (reply *actions.GetApplicationByIdResponse, err error) {
	id := req.GetId()

	store := s.manager.GetApplicationStore()
	app, err := store.GetByID(id)
	if err != nil {
		return
	}

	reply = &actions.GetApplicationByIdResponse{
		Exists: app != nil,
		App:    app.ToPb(),
	}

	return
}
