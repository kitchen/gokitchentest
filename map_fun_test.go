package main

// the basic jist of this test is for me to figure out some things I can and can't do with maps
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	HashMap map[string]string
	Number  int
}

// this is me knowing if I can `len()` an empty map, and what happens if I
// attempt to `range` an empty map.
func TestEmptyMap(t *testing.T) {
	foo := Foo{}

	length := len(foo.HashMap)
	assert.Equal(t, 0, length, "HashMap is empty")

	for k, v := range foo.HashMap {
		assert.Fail(t, fmt.Sprintf("we should not have gotten here: k: %s, v: %s", k, v))
	}
}

// this is just testing normal bits of a map, and access from outside of the ... thing
func TestMapWithStuff(t *testing.T) {
	foo := Foo{}
	foo.HashMap = make(map[string]string)
	foo.HashMap["bar"] = "baz"
	length := len(foo.HashMap)
	assert.Equal(t, 1, length, "HashMap should have one key/value pair")
	for k, v := range foo.HashMap {
		assert.Equal(t, "bar", k, "key should be bar")
		assert.Equal(t, "baz", v, "value should be baz")
	}
}
