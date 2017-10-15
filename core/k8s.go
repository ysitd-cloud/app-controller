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

func NewKubernetesDeployer(
	manager Manager,
	deployment appv1beta1.DeploymentInterface,
	ingress extv1beta1.IngressInterface,
	service corev1.ServiceInterface,
	secret corev1.SecretInterface,
) *KubernetesDeployer {
	return &KubernetesDeployer{
		deployment: deployment,
		manager:    manager,
		ingress:    ingress,
		service:    service,
		secret:     secret,
	}
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

func (deployer *KubernetesDeployer) CreateService(id string) error {
	template := deployer.createTemplate(id)
	service := template.GetService()
	_, err := deployer.service.Create(service)
	return err
}

func (deployer *KubernetesDeployer) UpdateService(id string) error {
	template := deployer.createTemplate(id)
	service := template.GetService()
	_, err := deployer.service.Update(service)
	return err
}

func (deployer *KubernetesDeployer) DeleteService(id string) error {
	template := deployer.createTemplate(id)
	return deployer.service.Delete(template.GetName(), &deletePolicy)
}

func (deployer *KubernetesDeployer) CreateSecret(id string) error {
	template := deployer.createTemplate(id)
	secret := template.GetSecret()
	_, err := deployer.secret.Create(secret)
	return err
}

func (deployer *KubernetesDeployer) UpdateSecret(id string) error {
	template := deployer.createTemplate(id)
	secret := template.GetSecret()
	_, err := deployer.secret.Update(secret)
	return err
}

func (deployer *KubernetesDeployer) DeleteSecret(id string) error {
	template := deployer.createTemplate(id)
	return deployer.secret.Delete(template.GetName(), &deletePolicy)
}

func (deployer *KubernetesDeployer) CreateIngress(id string) error {
	template := deployer.createTemplate(id)
	ingress := template.GetIngress()
	_, err := deployer.ingress.Create(ingress)
	return err
}

func (deployer *KubernetesDeployer) UpdateIngress(id string) error {
	template := deployer.createTemplate(id)
	ingress := template.GetIngress()
	_, err := deployer.ingress.Update(ingress)
	return err
}

func (deployer *KubernetesDeployer) DeleteIngress(id string) error {
	template := deployer.createTemplate(id)
	return deployer.ingress.Delete(template.GetName(), &deletePolicy)
}
