package core

import (
	"github.com/ysitd-cloud/app-controller/template/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appv1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

var propagationPolicy metav1.DeletionPropagation

var deletePolicy metav1.DeleteOptions = metav1.DeleteOptions{
	PropagationPolicy: &propagationPolicy,
}

type KubernetesDeployer struct {
	deployment appv1beta1.DeploymentInterface
	ingress    extv1beta1.IngressInterface
	service    corev1.ServiceInterface
	secret     corev1.SecretInterface
	manager    Manager
}

func (deployer *KubernetesDeployer) createTemplate(id string) kubernetes.KubernetesDeployment {
	app := deployer.manager.GetApplication(id)
	template := kubernetes.NewKubernetesDeployment(app)
	return template
}

func (deployer *KubernetesDeployer) Close() {
	deployer.manager.Close()
}

func (deployer *KubernetesDeployer) CreateDeployment(id string) error {
	template := deployer.createTemplate(id)
	deployment := template.GetDeployment()
	_, err := deployer.deployment.Create(deployment)
	return err
}

func (deployer *KubernetesDeployer) UpdateDeployment(id string) error {
	template := deployer.createTemplate(id)
	deployment := template.GetDeployment()
	_, err := deployer.deployment.Update(deployment)
	return err
}

func (deployer *KubernetesDeployer) DeleteDeployment(id string) error {
	template := deployer.createTemplate(id)
	return deployer.deployment.Delete(template.GetName(), &deletePolicy)
}
