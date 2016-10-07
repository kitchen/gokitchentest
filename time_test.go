package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeParse(t *testing.T) {
	foo, err := time.Parse(time.RFC3339, "2016-10-05T08:25:25-07:00")
	assert.NoError(t, err, "should have parsed successfully")
	assert.NotNil(t, foo, "should have parsed something")
}
