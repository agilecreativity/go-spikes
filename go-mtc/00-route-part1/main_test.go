package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Routing(t *testing.T) {
	a := assert.New(t)
	ts := httptest.NewServer(NewMux())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	a.NoError(err)
	body, err := ioutil.ReadAll(res.Body)
	a.Equal(string(body), "Users Index")
}
