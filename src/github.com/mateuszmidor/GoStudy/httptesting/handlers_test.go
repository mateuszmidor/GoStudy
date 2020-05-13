package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkStatus(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Got %q, expected %q", http.StatusText(actual), http.StatusText(expected))
	}
}
func TestAbout(t *testing.T) {
	// given
	r := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()

	// when
	HandleAbout(w, r)

	// then
	checkStatus(http.StatusOK, w.Code, t)
}

func TestAdmin(t *testing.T) {
	// given
	r := httptest.NewRequest("GET", "/admin", nil)
	w := httptest.NewRecorder()

	// when
	HandleAdmin(w, r)

	// then
	checkStatus(http.StatusForbidden, w.Code, t)
}

// This is like integration test; tests the multiplexer configuration + individual handlers
func TestMyServer(t *testing.T) {
	requestResponse := map[string]int{
		"/about":       http.StatusOK,
		"/admin":       http.StatusForbidden,
		"/nothinghere": http.StatusNotFound,
	}

	srv := NewMyServer()

	for request, expectedResponse := range requestResponse {

		t.Run(request[1:], func(t *testing.T) {
			// given
			r := httptest.NewRequest("GET", request, nil)
			w := httptest.NewRecorder()

			// when
			srv.ServeHTTP(w, r)

			// then
			checkStatus(expectedResponse, w.Code, t)
		})
	}
}
