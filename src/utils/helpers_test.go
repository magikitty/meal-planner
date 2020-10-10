package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpers_ConvertStringToInt(t *testing.T) {
	intExpected := 32
	intFromStringSucceess := ConvertStringToInt("32")

	assert.Equal(t, intExpected, intFromStringSucceess, "Failed to convert number string to int")
}
