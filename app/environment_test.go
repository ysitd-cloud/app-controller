package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironmentGetKeys(t *testing.T) {
	env := make(Environment)

	env["foo"] = "bar"

	assert.EqualValues(t, []string{"foo"}, env.GetKeys())
}
