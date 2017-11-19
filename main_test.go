package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRouter = getRouter()

func TestDefaultHandler(t *testing.T) {
	server := httptest.NewServer(testRouter)
	defer server.Close()
	resp, err := http.Get(server.URL + "/")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Buddy")
	assert.Contains(t, string(body), "Boy")
}

func TestHelloHandler(t *testing.T) {
	server := httptest.NewServer(testRouter)
	defer server.Close()
	resp, err := http.Get(server.URL + "/hello/friend")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Hi FRIEND")
}

func TestHandlerWithError(t *testing.T) {
	server := httptest.NewServer(testRouter)
	defer server.Close()
	resp, err := http.Get(server.URL + "/broken/handler")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), `html/template: "templates/missing.html" is undefined`)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestNewTemplateWithMissingTemplate(t *testing.T) {
	assert.Panics(t, func() { newTemplate("missing", nil, nil) })
}

func TestNewTemplateWithInvalidTemplate(t *testing.T) {
	assert.Panics(t, func() { newTemplate("invalid.html", nil, nil) })
}
