package main

import "encoding/binary"
import "testing"
import "bytes"
import "github.com/stretchr/testify/assert"

func TestUIntBytesThing(t *testing.T) {
	buf := new(bytes.Buffer)
	fooslice := make([]byte, 4)
	binary.BigEndian.PutUint32(fooslice, 42)
	written, err := buf.Write(fooslice)
	assert.NoError(t, err, "this should write successfully")
	assert.Equal(t, 4, written, "should write 4 bytes")
}
