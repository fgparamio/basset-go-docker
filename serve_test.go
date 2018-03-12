package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	assert.NotEmpty(t, getDataResponse(res).Users)
}

func TestGetUserWithPageOk(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"?page=1", nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, getDataResponse(res).Users)
}

func TestGetUserWrongParametersReturnFirstPage(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"?foo=1", nil)
	res, err := http.DefaultClient.Do(request)
	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, getDataResponse(res).Users)

}

func TestGetUserNotExistPageReturnEmpty(t *testing.T) {

	request, err := http.NewRequest("GET", usersUrl+"?page=999999999", nil)
	res, err := http.DefaultClient.Do(request)

	checkResponseCode(t, err, http.StatusOK, res.StatusCode)
	assert.Empty(t, getDataResponse(res).Users)

}

func TestBadUrlIsNotFound(t *testing.T) {

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

func getDataResponse(res *http.Response) main.PageDTO {
	body, _ := ioutil.ReadAll(res.Body)
	var page main.PageDTO
	json.Unmarshal([]byte(body), &page)
	return page
}
