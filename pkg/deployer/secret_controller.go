package deployer

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (sc *secretController) Create(secret *v1.Secret) (*v1.Secret, error) {
	return sc.client.Create(secret)
}

func (sc *secretController) Get(name string) (*v1.Secret, error) {
	return sc.client.Get(name, metav1.GetOptions{})
}

func (sc *secretController) Update(name string, env map[string][]byte) (*v1.Secret, error) {
	secret, err := sc.Get(name)
	if err != nil {
		return nil, err
	}

	secret.Data = env

	return sc.client.Update(secret)
}

func (sc *secretController) Delete(name string) error {
	return sc.client.Delete(name, deleteOptions)
}
