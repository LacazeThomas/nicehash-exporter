package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSecureTokenPositiveLen(t *testing.T) {
	str, err := GenerateSecureToken(10)
	assert.NoError(t, err)
	assert.Len(t, str, 10)
}

func TestGenerateSecureTokenNullableLen(t *testing.T) {
	str, err := GenerateSecureToken(0)
	assert.NoError(t, err)
	assert.Equal(t, str, "")
}
