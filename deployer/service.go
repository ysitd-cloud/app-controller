package deployer

import (
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/api/core/v1"
)

func (c *Controller) serviceClient() corev1.ServiceInterface {
	return c.client.CoreV1().Services(c.namespace)
}

func (c *Controller) CreateService(service *v1.Service) (*v1.Service, error) {
	return c.serviceClient().Create(service)
}

func (c *Controller) DeleteService(name string) error {
	return c.serviceClient().Delete(name, deleteOptions)
}
