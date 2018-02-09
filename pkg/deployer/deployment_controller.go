package deployer

import (
	"fmt"

	"k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *deploymentController) Create(deployment *v1beta2.Deployment) (*v1beta2.Deployment, error) {
	return c.client.Create(deployment)
}

func (c *deploymentController) Delete(name string) error {
	return c.client.Delete(name, deleteOptions)
}

func (c *deploymentController) Get(name string) (*v1beta2.Deployment, error) {
	return c.client.Get(name, metav1.GetOptions{})
}

func (c *deploymentController) UpdateImage(name, image, tag string) (*v1beta2.Deployment, error) {
	deployment, err := c.Get(name)
	if err != nil {
		return nil, err
	}

	deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", image, tag)

	return c.client.Update(deployment)
}
