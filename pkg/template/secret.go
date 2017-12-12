package template

import "k8s.io/api/core/v1"

func GenerateSecret(id string, env map[string]string) *v1.Secret {
	return &v1.Secret{
		ObjectMeta: getObjectMeta(id),
		Data:       convertToBytesMap(env),
		Type:       v1.SecretTypeOpaque,
	}
}

func convertToBytesMap(env map[string]string) map[string][]byte {
	bytes := make(map[string][]byte)
	for k, v := range env {
		bytes[k] = []byte(v)
	}

	return bytes
}
