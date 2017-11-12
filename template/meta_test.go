package template

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	id := "123"
	name := getName(id)
	assert.Equal(t, "app-123", name)
}

func TestGetObjectMeta(t *testing.T) {
	meta := getObjectMeta("123")
	assert.Equal(t, "app-123", meta.Name)
}
