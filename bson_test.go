package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
)

func TestBsonPlayground(t *testing.T) {
	foo := bson.M{"a": 1, "b": 2}
	foo["bar"] = "baz"
	assert.NotNil(t, foo, "noop")

	bar := bson.M{"foo": time.Now()}
	t.Logf("bson time: %v", bar)
}
