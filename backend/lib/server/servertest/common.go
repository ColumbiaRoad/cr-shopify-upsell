package servertest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
)

// Customizer is a type for request-modifying functions that can be used to
// customize requests created by the shorthand request functions (Do, Get, Post, Delete)
type Customizer func(*http.Request)

// Do performs given method on url with handler and returns an
// httptest.ResponseRecorder to inspect the resulting response
func Do(handler http.Handler, method, url string, v interface{}, fs ...Customizer) *httptest.ResponseRecorder {
	var body bytes.Buffer
	if v != nil {
		if err := json.NewEncoder(&body).Encode(&v); err != nil {
			log.Fatalf("encode json: %v", err)
		}
	}
	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		log.Fatalf("create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	for _, f := range fs {
		f(req)
	}

	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	return res
}

// Post is a shorthand for Do(handler, "POST", url, v)
func Post(handler http.Handler, url string, v interface{}, fs ...Customizer) *httptest.ResponseRecorder {
	return Do(handler, "POST", url, v, fs...)
}

// Get is a shorthand for Do(handler, "GET", url, nil)
func Get(handler http.Handler, url string, fs ...Customizer) *httptest.ResponseRecorder {
	return Do(handler, "GET", url, nil, fs...)
}

// Delete is a shorthand for Do(handler, "DELETE", url, nil)
func Delete(handler http.Handler, url string) *httptest.ResponseRecorder {
	return Do(handler, "DELETE", url, nil)
}

// AddHeader is a request customizer that adds given header to the request
func AddHeader(name, value string) Customizer {
	return func(r *http.Request) {
		r.Header.Add(name, value)
	}
}
