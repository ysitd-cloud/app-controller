package template

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDeployment(t *testing.T) {
	env := map[string]string {
		"foo": "bar",
	}
	deployment := GenerateDeployment("123", "golang", "1.9-alpine", env)
	assert := assert.New(t)

	spec := deployment.Spec
	assert.Equal(int32(1), *spec.Replicas)

	container := spec.Template.Spec.Containers[0]
	assert.Equal("golang:1.9-alpine", container.Image)
}
