package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTime(t *testing.T) {
	time, err := GetTime()
	assert.NoError(t, err)
	assert.NotEmpty(t, time)
}
