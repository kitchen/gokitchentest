package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type foo int

const fooThing = 0

type bar int

const barThing = 0

func TestUnderlyingTypes(t *testing.T) {
	assert.Exactly(t, fooThing, barThing, "foo and bar are not the same")
}
