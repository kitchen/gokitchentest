package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteSliceConcatenation(t *testing.T) {
	var foo []byte
	foo = append(foo, []byte{0x01, 0x02}...)
	foo = append(foo, []byte("foo")...)
	foo = append(foo, []byte{0x01, 0x02}...)
	assert.Equal(t, []byte{0x01, 0x02, 0x66, 0x6f, 0x6f, 0x01, 0x02}, foo, "should have the same bytes")
}
