package deployer

import (
	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
)

type Controller interface {
	SetClient(client kubernetes.Interface)
	SetNamespace(namespace string)

	CreateDeployment(deployment *v1beta2.Deployment) (*v1beta2.Deployment, error)
	DeleteDeployment(name string) error
	GetDeployment(name string) (*v1beta2.Deployment, error)
	UpdateDeploymentImage(name, image, tag string) (*v1beta2.Deployment, error)

	CreateIngress(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error)
	DeleteIngress(name string) error
	GetIngress(name string) (*v1beta1.Ingress, error)
	UpdateIngressDomain(name, domain string) (*v1beta1.Ingress, error)

	CreateSecret(secret *v1.Secret) (*v1.Secret, error)
	DeleteSecret(name string) error
	GetSecret(name string) (*v1.Secret, error)
	UpdateSecret(name string, env map[string][]byte) (*v1.Secret, error)

	CreateService(service *v1.Service) (*v1.Service, error)
	DeleteService(name string) error
}

type controller struct {
	client    kubernetes.Interface
	namespace string
}
