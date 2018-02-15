package deployer

import (
	"k8s.io/api/core/v1"
)

func (sc *serviceController) Create(service *v1.Service) (*v1.Service, error) {
	return sc.client.Create(service)
}

func (sc *serviceController) Delete(name string) error {
	return sc.client.Delete(name, deleteOptions)
}
