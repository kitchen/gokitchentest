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
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, foobyteswritable, "starts off with 4 0 bytes")
	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobyteswritable, uint32(len(foobytes)))
	}, "aha!")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x08}, foobyteswritable, "should be 8")

}

func TestPutUint32Again(t *testing.T) {
	foobytes := make([]byte, 4)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00}, foobytes, "starts off with 4 0 bytes")
	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobytes, 42)
	}, "it should work")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x2a}, foobytes, "should be 42")

	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobytes, 43)
	}, "it should work")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x2b}, foobytes, "should be 43")

	// ok, so I want to put ANOTHER uint32 on this, can I do that?
	foobytes = append(foobytes, []byte{0x00, 0x00, 0x00, 0x00}...)
	assert.NotPanics(t, func() {
		binary.BigEndian.PutUint32(foobytes, 44)
	}, "it should work")
	assert.NotEqual(t, []byte{0x00, 0x00, 0x00, 0x2b, 0x00, 0x00, 0x00, 0x2c}, foobytes, "it is not 43, followed by 44")
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x00}, foobytes, "it is in fact 44, followed by 0")
	// seems PutUint32 replaces the first 4 bytes of a byte slice. cool!

}
