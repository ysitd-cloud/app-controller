package template

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/api/core/v1"
	"testing"
)

func TestGenerateService(t *testing.T) {
	service := GenerateService("123")
	assert := assert.New(t)

	spec := service.Spec

	assert.Equal(v1.ServiceTypeClusterIP, spec.Type)

	port := spec.Ports[0]
	assert.Equal(v1.ProtocolTCP, port.Protocol)
}
