package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func returnBool(ret bool) bool {
	return ret
}

func TestBoolFunctionArgument(t *testing.T) {
	assert.True(t, returnBool(true == true), "test that we can use an expression as an argument to a function")
}
