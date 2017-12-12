package template

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	"testing"
)

func TestGenerateSecret(t *testing.T) {
	envs := map[string]string{
		"foo": "bar",
	}
	secret := GenerateSecret("123", envs)

	assert := assert.New(t)

	assert.Equal(v1.SecretTypeOpaque, secret.Type)
	assert.EqualValues([]byte("bar"), secret.Data["foo"])
}
