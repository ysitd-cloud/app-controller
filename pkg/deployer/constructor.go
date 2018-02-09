package deployer

import "k8s.io/client-go/kubernetes"

func NewController(client kubernetes.Interface, namespace string) Controller {
	provider := &k8sClientProvider{
		client:    client,
		namespace: namespace,
	}
	return &controller{
		deployment: newDeploymentController(provider),
		service:    newServiceController(provider),
		secret:     newSecretController(provider),
		ingress:    newIngressController(provider),
	}
}

func newDeploymentController(provider clientProvider) DeploymentController {
	c := new(deploymentController)
	c.client = provider.deploymentClient()
	return c
}

func newServiceController(provider clientProvider) ServiceController {
	c := new(serviceController)
	c.client = provider.serviceClient()
	return c
}

func newSecretController(provider clientProvider) SecretController {
	c := new(secretController)
	c.client = provider.secretClient()
	return c
}

func newIngressController(provider clientProvider) IngressController {
	c := new(ingressController)
	c.client = provider.ingressClient()
	return c
}
