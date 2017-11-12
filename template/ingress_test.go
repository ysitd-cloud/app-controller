package template

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateIngress(t *testing.T) {
	ingress := GenerateIngress("123", "example.com")
	rule := ingress.Spec.Rules[0]
	assert := assert.New(t)

	assert.Equal("example.com", rule.Host)

	service := rule.HTTP.Paths[0].Backend.ServiceName
	assert.Equal("app-123", service)
}
