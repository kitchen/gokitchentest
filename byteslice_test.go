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

type EmptyByteSliceThing struct {
	ByteSlice []byte
}

func TestStructedByteSlice(t *testing.T) {
	foo := EmptyByteSliceThing{}
	assert.Nil(t, foo.ByteSlice, "it hasn't been initialized yet, so it should be nil")

	foo.ByteSlice = make([]byte, 16)
	assert.NotNil(t, foo.ByteSlice, "it has been initialized, so it shouldn't be nil")
}

// func TestMapByteSliceByteSlice(t *testing.T) {
// 	foo := map[[]byte][]byte{
// 		[]byte("foo"): []byte("bar"),
// 	}
//
// 	assert.True(t, true, "this doesn't work")
// }

// WIARTTD: map[[]byte][]byte
func TestSliceOfSlicesOfByteSlices(t *testing.T) {
	foo := [][][]byte{
		[][]byte{
			[]byte("foo"), []byte("bar"),
		},
		[][]byte{
			[]byte("foo1"), []byte("bar1"),
		},
	}
	assert.NotNil(t, foo, "it's not nil!")
	assert.True(t, true, "this works!")
}

func TestSliceSlice(t *testing.T) {
	foo := make([]byte, 1200)
	total := 0

	for {
		if len(foo) == 0 {
			break
		}
		cur := foo[:100]
		foo = foo[100:]
		total += len(cur)
	}

	assert.Equal(t, 1200, total, "we should have chunked 1200 bytes off")

	foo = make([]byte, 0, 10)
	foo = append(foo, []byte("abcde12345")...)

	chunksize := 3
	for {
		if len(foo) == 0 {
			break
		}

		if chunksize > len(foo) {
			chunksize = len(foo)
		}

		cur := foo[:chunksize]
		foo = foo[chunksize:]
		t.Logf("cur: %s, foo: %s", string(cur), string(foo))
	}

}
