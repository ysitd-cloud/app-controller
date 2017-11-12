package deployer

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"k8s.io/client-go/pkg/apis/apps/v1beta1"
)

func (c *Controller) deploymentClient() appv1beta1.DeploymentInterface {
	return c.client.AppsV1beta1().Deployments(c.namespace)
}

func (c *Controller) CreateDeployment(deployment *v1beta1.Deployment) (*v1beta1.Deployment, error) {
	return c.deploymentClient().Create(deployment)
}

func (c *Controller) DeleteDeployment(name string) error {
	return c.deploymentClient().Delete(name, deleteOptions)
}

func (c *Controller) GetDeployment(name string) (*v1beta1.Deployment, error) {
	return c.deploymentClient().Get(name, metav1.GetOptions{})
}

func (c *Controller) UpdateDeploymentImage(name, image, tag string) (*v1beta1.Deployment, error) {
	deployment, err := c.GetDeployment(name)
	if err != nil {
		return nil, err
	}

	deployment.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", image, tag)

	return c.deploymentClient().Update(deployment)
}
