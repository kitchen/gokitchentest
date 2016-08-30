package main

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomBytes(t *testing.T) {
	bytes := make([]byte, 8)
	n, err := rand.Read(bytes)
	assert.NoError(t, err, "this should not error")
	assert.Equal(t, 8, n, "n should equal the size of the byte slice")

}
