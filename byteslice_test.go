package main

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyByteSlice(t *testing.T) {
	bytes := make([]byte, 8)
	assert.NotNil(t, bytes, "a fresh byte slice is nil")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, bytes, "fresh byte slice is full of zeros")

	n, err := rand.Read(bytes)
	assert.NoError(t, err, "this should not error")
	assert.Equal(t, 8, n, "n should equal the size of the byte slice")

	assert.NotNil(t, bytes, "a non-empty byte slice is not nil")
}
