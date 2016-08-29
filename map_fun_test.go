package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	HashMap map[string]string
	Number  int
}

func TestEmptyMap(t *testing.T) {
	foo := Foo{}

	length := len(foo.HashMap)
	assert.Equal(t, 0, length, "HashMap is empty")
	for k, v := range foo.HashMap {
		assert.Fail(t, fmt.Sprintf("we should not have gotten here: k: %s, v: %s", k, v))
	}
}

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
