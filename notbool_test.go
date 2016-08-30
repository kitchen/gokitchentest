package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotBool(t *testing.T) {
	if !true {
		assert.Fail(t, "true isn't true")
	}
}
