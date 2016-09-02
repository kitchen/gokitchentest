package main

import (
	"bytes"
	"compress/gzip"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGzipWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	writer := gzip.NewWriter(buf)
	_, err := writer.Write([]byte("foo"))
	assert.NoError(t, err, "successfully written")
	err = writer.Close()
	assert.NoError(t, err, "successfully closed")
	assert.NotEqual(t, []byte{}, buf.Bytes(), "stuff was written to the buffer")
	oldlen := len(buf.Bytes())

	newbuf := new(bytes.Buffer)
	writer.Reset(newbuf)
	_, err = writer.Write([]byte("bar"))
	assert.NoError(t, err, "successfully written")
	err = writer.Close()
	assert.NoError(t, err, "successfully closed")
	assert.NotEqual(t, []byte{}, newbuf.Bytes(), "stuff was written to the new buffer")
	assert.Equal(t, oldlen, len(buf.Bytes()), "nothing was written to the old buffer")
}
