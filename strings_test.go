package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSplit(t *testing.T) {
	foo := strings.Split("abc - def", " - ")
	assert.Equal(t, "abc", foo[0], "should be abc")
	assert.Equal(t, "def", foo[1], "should be def")

	assert.Equal(t, 2, len(foo), "should have returned 2 things")
}
