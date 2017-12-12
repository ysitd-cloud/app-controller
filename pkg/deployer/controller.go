package deployer

import (
	"fmt"

	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	appv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionV1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func (c *controller) SetClient(client kubernetes.Interface) {
	c.client = client
}

func (c *controller) SetNamespace(namespace string) {
	c.namespace = namespace
}

func (c *controller) deploymentClient() appv1beta2.DeploymentInterface {
	return c.client.AppsV1beta2().Deployments(c.namespace)
}

func (c *controller) CreateDeployment(deployment *v1beta2.Deployment) (*v1beta2.Deployment, error) {
	return c.deploymentClient().Create(deployment)
}

func (c *controller) DeleteDeployment(name string) error {
	return c.deploymentClient().Delete(name, deleteOptions)
}

func (c *controller) GetDeployment(name string) (*v1beta2.Deployment, error) {
	return c.deploymentClient().Get(name, metav1.GetOptions{})
}

func (c *controller) UpdateDeploymentImage(name, image, tag string) (*v1beta2.Deployment, error) {
	deployment, err := c.GetDeployment(name)
	if err != nil {
		return nil, err
	}

	deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", image, tag)

	return c.deploymentClient().Update(deployment)
}

func (c *controller) ingressClient() extensionV1beta1.IngressInterface {
	return c.client.ExtensionsV1beta1().Ingresses(c.namespace)
}

func (c *controller) CreateIngress(ingress *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	return c.ingressClient().Create(ingress)
}

func (c *controller) DeleteIngress(name string) error {
	return c.ingressClient().Delete(name, deleteOptions)
}

func (c *controller) GetIngress(name string) (*v1beta1.Ingress, error) {
	return c.ingressClient().Get(name, metav1.GetOptions{})
}

func (c *controller) UpdateIngressDomain(name, domain string) (*v1beta1.Ingress, error) {
	ingress, err := c.GetIngress(name)
	if err != nil {
		return nil, err
	}

	ingress.Spec.Rules[0].Host = domain

	return c.ingressClient().Update(ingress)
}

func (c *controller) secretClient() corev1.SecretInterface {
	return c.client.CoreV1().Secrets(c.namespace)
}

func (c *controller) CreateSecret(secret *v1.Secret) (*v1.Secret, error) {
	return c.secretClient().Create(secret)
}

func (c *controller) DeleteSecret(name string) error {
	return c.secretClient().Delete(name, deleteOptions)
}

func (c *controller) GetSecret(name string) (*v1.Secret, error) {
	return c.secretClient().Get(name, metav1.GetOptions{})
}

func (c *controller) UpdateSecret(name string, env map[string][]byte) (*v1.Secret, error) {
	secret, err := c.GetSecret(name)
	if err != nil {
		return nil, err
	}

	secret.Data = env

	return c.secretClient().Update(secret)
}

func (c *controller) serviceClient() corev1.ServiceInterface {
	return c.client.CoreV1().Services(c.namespace)
}

func (c *controller) CreateService(service *v1.Service) (*v1.Service, error) {
	return c.serviceClient().Create(service)
}

func (c *controller) DeleteService(name string) error {
	return c.serviceClient().Delete(name, deleteOptions)
}
