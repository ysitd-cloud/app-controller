package deployer

func (c *controller) GetDeploymentController() DeploymentController {
	return c.deployment
}

func (c *controller) GetIngressController() IngressController {
	return c.ingress
}

func (c *controller) GetSecretController() SecretController {
	return c.secret
}

func (c *controller) GetServiceController() ServiceController {
	return c.service
}
