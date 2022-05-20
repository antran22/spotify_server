package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeTestRequest(t testing.TB, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	return resp, string(respBody)
}

func guardIntegrationTest(t testing.TB) {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping integration test")
	}
}

func getTestServer() *httptest.Server {
	router := getMainRouter()
	ts := httptest.NewServer(router)
	return ts
}
