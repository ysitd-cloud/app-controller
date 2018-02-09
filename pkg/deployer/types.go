package deployer

import (
	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

type Controller interface {
	GetDeploymentController() DeploymentController
	GetIngressController() IngressController
	GetSecretController() SecretController
	GetServiceController() ServiceController
}

type DeploymentController interface {
	Create(deployment *v1beta2.Deployment) (*v1beta2.Deployment, error)
	Delete(name string) error
	Get(name string) (*v1beta2.Deployment, error)
	UpdateImage(name, image, tag string) (*v1beta2.Deployment, error)
}

type IngressController interface {
	Create(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error)
	Get(name string) (*v1beta1.Ingress, error)
	UpdateDomain(name, domain string) (*v1beta1.Ingress, error)
	Delete(name string) error
}

type SecretController interface {
	Create(secret *v1.Secret) (*v1.Secret, error)
	Get(name string) (*v1.Secret, error)
	Update(name string, env map[string][]byte) (*v1.Secret, error)
	Delete(name string) error
}

type ServiceController interface {
	Create(service *v1.Service) (*v1.Service, error)
	Delete(name string) error
}
