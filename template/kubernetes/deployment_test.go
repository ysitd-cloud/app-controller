package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeployment(t *testing.T) {
	asserts := assert.New(t)

	app := createTestApplication()
	deployment := NewKubernetesDeployment(app)
	appDeployment := deployment.GetDeployment()

	asserts.EqualValues(32, *appDeployment.Spec.Replicas, "Wrong replicas")

	container := appDeployment.Spec.Template.Spec.Containers[0]

	asserts.Equal("app-testing", container.Name, "Wrong Name")
	asserts.Equal("golang:1.9-alpine", container.Image, "Wrong Image")

	port := container.Ports[0]
	asserts.Equal("http", port.Name, "Wrong Port Name")
	asserts.Equal(int32(12345), port.ContainerPort, "Wrong Port Number")

	env := container.Env[0]

	asserts.Equal("usage", env.Name, "Wrong env key")

	value := env.ValueFrom.SecretKeyRef
	asserts.Equal("app-testing", value.Name, "Wrong env secret name")
	asserts.Equal("usage", value.Key, "Wrong secret key")
}
