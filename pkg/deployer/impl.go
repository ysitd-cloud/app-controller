package deployer

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

type controller struct {
	deployment DeploymentController
	ingress    IngressController
	secret     SecretController
	service    ServiceController
}

type k8sClientProvider struct {
	client    kubernetes.Interface
	namespace string
}

type clientProvider interface {
	SetClient(client kubernetes.Interface)
	getClient() kubernetes.Interface
	getNamespace() string
	deploymentClient() v1beta2.DeploymentInterface
	ingressClient() v1beta1.IngressInterface
	secretClient() v1.SecretInterface
	serviceClient() v1.ServiceInterface
}

type deploymentController struct {
	client v1beta2.DeploymentInterface
}

type ingressController struct {
	client v1beta1.IngressInterface
}

type secretController struct {
	client v1.SecretInterface
}

type serviceController struct {
	client v1.ServiceInterface
}
