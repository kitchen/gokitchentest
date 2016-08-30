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

func TestDifferentMakes(t *testing.T) {
	// ok, so the docs for make([]foo, 0, 10) a length 0, capacity 10
	// but that doesn't mean it's *max* capacity, just that it pre-allocates that much space.
	foo := make([]int, 0, 10)
	assert.Equal(t, 0, len(foo), "initial length is 0")

	for i := 0; i < 14; i++ {
		foo = append(foo, 42)
	}
	assert.Equal(t, 14, len(foo), "length is 14")

	// this creates an int slice of size 10 full of zeros
	bar := make([]int, 10)
	assert.Equal(t, 10, len(bar), "initial length is 10")
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, bar, "full of zeros")
	bar = append(bar, 42)
	assert.Equal(t, 11, len(bar), "it gets another elemented added to the end")
}
