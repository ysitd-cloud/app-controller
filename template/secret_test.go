package template

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/pkg/api/v1"
)

func TestGenerateSecret(t *testing.T) {
	envs := map[string]string {
		"foo": "bar",
	}
	secret := GenerateSecret("123", envs)

	assert := assert.New(t)

	assert.Equal(v1.SecretTypeOpaque, secret.Type)
	assert.EqualValues([]byte("bar"), secret.Data["foo"])
}
