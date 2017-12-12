package deployer

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	extensionV1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func (c *Controller) ingressClient() extensionV1beta1.IngressInterface {
	return c.client.ExtensionsV1beta1().Ingresses(c.namespace)
}

func (c *Controller) CreateIngress(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	return c.ingressClient().Create(ingress)
}

func (c *Controller) DeleteIngress(name string) error {
	return c.ingressClient().Delete(name, deleteOptions)
}

func (c *Controller) GetIngress(name string) (*v1beta1.Ingress, error) {
	return c.ingressClient().Get(name, v1.GetOptions{})
}

func (c *Controller) UpdateIngressDomain(name, domain string) (*v1beta1.Ingress, error) {
	ingress, err := c.GetIngress(name)
	if err != nil {
		return nil, err
	}

	ingress.Spec.Rules[0].Host = domain

	return c.ingressClient().Update(ingress)
}
