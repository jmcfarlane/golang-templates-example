package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRoute(t *testing.T) {
	server := httptest.NewServer(getRouter())
	defer server.Close()
	resp, err := http.Get(server.URL + "/")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Buddy")
	assert.Contains(t, string(body), "Boy")
}
