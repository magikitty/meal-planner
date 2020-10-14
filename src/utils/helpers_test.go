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

func TestHelpers_GetFileData(t *testing.T) {
	byteArrayExpected := []byte{0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa, 0x7d}
	byteArrayActual := GetFileData("../../data/tests/test-getfiledata.json")

	assert.Equal(t, byteArrayExpected, byteArrayActual, "Failed to get expected byte array from file")
}
