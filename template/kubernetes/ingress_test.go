package kubernetes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIngress(t *testing.T) {
	asserts := assert.New(t)

	app := createTestApplication()
	deployment := NewKubernetesDeployment(app)
	ingress := deployment.GetIngress()

	rule := ingress.Spec.Rules[0]
	asserts.Equal("test.app.ysitd.cloud", rule.Host, "Wrong hostname")

	path := rule.HTTP.Paths[0]
	asserts.Equal("/", path.Path, "Wrong path")
}
