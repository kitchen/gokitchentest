package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntNilTest struct {
	Foo int
}

func TestCanIntsBeNil(t *testing.T) {
	foo := IntNilTest{}
	assert.NotNil(t, foo.Foo, "int cannot be nil")
	assert.Equal(t, 0, foo.Foo, "it is initialized as 0")
}
