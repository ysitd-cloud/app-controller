package deployer

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/api/core/v1"
)

func (c *Controller) secretClient() corev1.SecretInterface {
	return c.client.Secrets(c.namespace)
}

func (c *Controller) CreateSecret(secret *v1.Secret) (*v1.Secret, error) {
	return c.secretClient().Create(secret)
}

func (c *Controller) DeleteSecret(name string) error {
	return c.secretClient().Delete(name, deleteOptions)
}

func (c *Controller) GetSecret(name string) (*v1.Secret, error) {
	return c.secretClient().Get(name, metav1.GetOptions{})
}

func (c *Controller) UpdateSecret(name string, env map[string][]byte) (*v1.Secret, error) {
	secret, err := c.GetSecret(name)
	if err != nil {
		return nil, err
	}

	secret.Data = env

	return c.secretClient().Update(secret)
}
