package deployer

import (
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *ingressController) Create(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	return c.client.Create(ingress)
}

func (c *ingressController) Delete(name string) error {
	return c.client.Delete(name, deleteOptions)
}

func (c *ingressController) Get(name string) (*v1beta1.Ingress, error) {
	return c.client.Get(name, metav1.GetOptions{})
}

func (c *ingressController) UpdateDomain(name, domain string) (*v1beta1.Ingress, error) {
	ingress, err := c.Get(name)
	if err != nil {
		return nil, err
	}

	ingress.Spec.Rules[0].Host = domain

	return c.client.Update(ingress)
}
