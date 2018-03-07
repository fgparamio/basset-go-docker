package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"."
	"github.com/stretchr/testify/assert"
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	usersUrl string
)

func init() {
	server = httptest.NewServer(main.Handlers())   //Creating new server with the user handlers
	usersUrl = fmt.Sprintf("%s/users", server.URL) //Grab the address for the API endpoint
}

func TestGetUserWithoutPageOk(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl, nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.NotNil(t, res.Body.Read)
}

func TestGetUserWithPageOk(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"?page=1", nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.NotNil(t, res.Body.Read)
}

func TestGetUserWrongParameters(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"?foo=1", nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.NotNil(t, res.Body.Read)
}

func TestBadUrlIs400(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"/foo", nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusNotFound, res.StatusCode)

}

func BenchmarkUserApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		request, _ := http.NewRequest("GET", usersUrl, nil)
		http.DefaultClient.Do(request)
	}
}

func checkResponseCode(t *testing.T, err error, expected, actual int) {
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
