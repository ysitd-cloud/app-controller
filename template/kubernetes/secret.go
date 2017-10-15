package kubernetes

import "k8s.io/client-go/pkg/api/v1"

func (d *deploymentV1) GetSecret() *v1.Secret {
	return &v1.Secret{
		ObjectMeta: d.getObjectMeta(),
		Data:       d.application.GetEnvironment().ToBytesMap(),
		Type:       v1.SecretTypeOpaque,
	}
}
