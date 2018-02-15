package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetName(t *testing.T) {
	id := "123"
	name := GetName(id)
	assert.Equal(t, "app-123", name)
}

func TestGetObjectMeta(t *testing.T) {
	meta := getObjectMeta("123")
	assert.Equal(t, "app-123", meta.Name)
}
