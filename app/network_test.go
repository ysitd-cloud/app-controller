package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNetwork(t *testing.T) {
	asserts := assert.New(t)

	domain := "foo.bar"
	var port int32 = 1234
	network := NewNetwork(domain, port)

	asserts.Equal(domain, network.GetDomain(), "Wrong domain")
	asserts.Equal(port, network.GetPort(), "Wrong port")
}
