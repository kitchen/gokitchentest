package main

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutUint32(t *testing.T) {
	foobytes := []byte("aoeuhtns")
	foobyteswritable := make([]byte, 0, 4+len(foobytes))
	assert.Equal(t, []byte{}, foobyteswritable, "slice starts off empty")

	assert.Panics(t, func() {
		binary.BigEndian.PutUint32(foobyteswritable, uint32(len(foobytes)))
	}, "this doesn't work")

	foobyteswritable = make([]byte, 0, 4)
	assert.Equal(t, []byte{}, foobyteswritable, "slice starts off empty")
	assert.Panics(t, func() {
		binary.BigEndian.PutUint32(foobyteswritable, uint32(len(foobytes)))
	}, "nor does this")

	foobyteswritable = make([]byte, 4)
	assert.NotEqual(t, []byte{}, foobyteswritable, "slice does not start off empty")
	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobyteswritable, uint32(len(foobytes)))
	}, "but this does")

	foobyteswritable = make([]byte, 4, 4+len(foobytes))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, foobyteswritable, "stars off with 4 0 bytes")
	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobyteswritable, uint32(len(foobytes)))
	}, "aha!")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x08}, foobyteswritable, "should be 8")

}
