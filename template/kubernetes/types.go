package kubernetes

import (
	"github.com/ysitd-cloud/app-controller/app"
	apps "k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

type KubernetesDeployment interface {
	GetApplication() app.Application
	GetSecret() *v1.Secret
	GetService() *v1.Service
	GetIngress() *v1beta1.Ingress
	GetDeployment() *apps.Deployment
}

func NewKubernetesDeployment(application app.Application) KubernetesDeployment {
	return &deploymentV1{
		application: application,
	}
}

type deploymentV1 struct {
	application app.Application
}
