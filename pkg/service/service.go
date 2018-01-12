package service

import (
	"code.ysitd.cloud/component/deployer/pkg/manager"
	"code.ysitd.cloud/component/deployer/pkg/template"
	"github.com/ysitd-cloud/grpc-schema/deployer/actions"
	"github.com/ysitd-cloud/grpc-schema/deployer/models"
	"golang.org/x/net/context"
)

func (s *service) ListApplicationsByUsername(_ context.Context, req *actions.ListApplicationsByUsernameRequest) (reply *actions.ListApplicationsByUsernameReply, err error) {
	username := req.GetUsername()
	apps, err := s.manager.GetApplicationByOwner(username)
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
	confirm, e, err := s.manager.CreateApplication(app)
	if err != nil {
		return
	}
	defer close(confirm)

	env := app.Environment
	secret := template.GenerateSecret(app.ID, env)
	if _, err = s.deployer.CreateSecret(secret); err != nil {
		confirm <- false
		return
	}

	d := app.Deployment
	deployment := template.GenerateDeployment(app.ID, d.Image, d.Tag, env)
	if _, err = s.deployer.CreateDeployment(deployment); err != nil {
		confirm <- false
		return
	}

	service := template.GenerateService(app.ID)
	if _, err = s.deployer.CreateService(service); err != nil {
		confirm <- false
		return
	}

	ingress := template.GenerateIngress(app.ID, app.Network.GetDomain())
	if _, err = s.deployer.CreateIngress(ingress); err != nil {
		confirm <- false
		return
	}

	confirm <- true
	err = <-e

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

	confirm, e, err := s.manager.UpdateDeployment(id, deployment)
	if err != nil {
		return
	}
	defer close(confirm)

	_, err = s.deployer.UpdateDeploymentImage(template.GetName(id), deployment.GetImage(), deployment.GetTag())
	if err != nil {
		confirm <- false
		return
	}

	confirm <- true
	err = <-e
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

	app, err := s.manager.GetApplicationByID(id)
	if err != nil {
		return
	}

	reply = &actions.GetApplicationByIdResponse{
		Exists: app != nil,
		App:    app.ToPb(),
	}

	return
}
