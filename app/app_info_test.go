package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetaInformation(t *testing.T) {

	asserts := assert.New(t)

	expectImage := "go"
	expectTag := "1.9-alpine"

	info := NewMetaInformation(expectImage, expectTag)

	asserts.Equal(expectImage, info.GetImage())
	asserts.Equal(expectTag, info.GetTag())
}
