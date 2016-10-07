package main

import (
	"fmt"
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

func TestStringIntMapNoKey(t *testing.T) {
	foo := make(map[string]int)
	bar := foo["bar"]
	if bar == 0 {
		t.Logf("foobar")
	}
	foo["baz"]++
	assert.Equal(t, 1, foo["baz"], "foo[baz] should be 1")
}

const Ten = 10

func TestNegativeConstant(t *testing.T) {
	assert.Equal(t, -10, -Ten, "-Constant works")
}

const FooInt = 42

func TestPrintf(t *testing.T) {

	assert.Equal(t, "42", fmt.Sprintf("%d", FooInt), "should be 42")
}
