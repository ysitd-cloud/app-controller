package deployer

import (
	"k8s.io/client-go/kubernetes"
	appv1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	extensionV1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

func (c *k8sClientProvider) SetClient(client kubernetes.Interface) {
	c.client = client
}

func (c *k8sClientProvider) SetNamespace(namespace string) {
	c.namespace = namespace
}

func (c *k8sClientProvider) getClient() kubernetes.Interface {
	return c.client
}

func (c *k8sClientProvider) getNamespace() string {
	return c.namespace
}

func (c *k8sClientProvider) deploymentClient() appv1beta2.DeploymentInterface {
	return c.client.AppsV1beta2().Deployments(c.namespace)
}

func (c *k8sClientProvider) ingressClient() extensionV1beta1.IngressInterface {
	return c.client.ExtensionsV1beta1().Ingresses(c.namespace)
}

func (c *k8sClientProvider) secretClient() corev1.SecretInterface {
	return c.client.CoreV1().Secrets(c.namespace)
}

func (c *k8sClientProvider) serviceClient() corev1.ServiceInterface {
	return c.client.CoreV1().Services(c.namespace)
}
