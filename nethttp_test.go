package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/foo/bar?baz=bat&multi=123&multi=456", new(bytes.Buffer))
	assert.NoError(t, err, "there wasn't an error")
	err = req.ParseForm()
	assert.NoError(t, err, "we successfully parsed the form")

	err = req.ParseForm()
	assert.NoError(t, err, "we can do this multiple times and not care")
	t.Logf("form values: %v", req.Form)

	t.Logf("baz value: %v", req.Form.Get("baz"))
	t.Logf("multi value: %v", req.Form.Get("multi"))
	t.Logf("non-existent value: %v", req.Form.Get("aoeuhtns"))

	if foo := req.Form.Get("baz"); foo != "" {
		t.Logf("neat!")
	}

	if foo := req.Form.Get("aoeuhtns"); foo != "" {
		t.Logf("not neat!")
	} else {
		t.Logf("double neat!")
	}
}
