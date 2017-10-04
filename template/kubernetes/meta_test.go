package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetApplication(t *testing.T) {
	app := createTestApplication()

	deployment := NewKubernetesDeployment(app)

	assert.Equal(t, app, deployment.GetApplication())
}
