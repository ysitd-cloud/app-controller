package kubernetes

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/pkg/api/v1"
	"testing"
)

func TestGetSecret(t *testing.T) {
	asserts := assert.New(t)

	app := createTestApplication()
	deployment := NewKubernetesDeployment(app)
	secret := deployment.GetSecret()

	asserts.Equal("app-testing", secret.Name)
	asserts.EqualValues([]byte("test"), secret.Data["usage"])
	asserts.Equal(v1.SecretTypeOpaque, secret.Type)
}
