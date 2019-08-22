package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

// execute an HTTP request and return the response
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

// check HTTP response code actual against expected
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code is %d. Got %d", expected, actual)
	}
}

// test that full application runs
func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

// test that health status works
func TestHealthStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

// test default name response on hello path
func TestDefaultNameHello(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if response.Body.String() != "Hello from Go, Roald!" {
		t.Errorf("Expected 'Hello from Go, Roald!', but received %s", response.Body.String())
	}
}

// test a valid name on the hello path
func TestValidNameHello(t *testing.T) {
	name := "Scott"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/hello/%s", name), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if response.Body.String() != fmt.Sprintf("Hello from Go, %s!", name) {
		t.Errorf("Expected 'Hello from Go, %s!', but received %s", name, response.Body.String())
	}
}

// test an invalid name on the hello path
func TestInvalidNameHello(t *testing.T) {
	name := "1234"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/hello/%s", name), nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

