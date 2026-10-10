package add

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// resultBody mirrors api.ResultSuccess for decoding success responses.
type resultBody struct {
	Result float64 `json:"result"`
}

// errorBody mirrors api.ResultError for decoding error responses.
type errorBody struct {
	Error string `json:"error"`
}

func floatptr(f float64) *float64 { return &f }

// TestAddHandler exercises the handler logic directly, table-driven,
// covering valid sums and every parameter-validation failure mode.
func TestAddHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantCode   int
		wantResult *float64
		wantError  string
	}{
		{
			name:       "positive integers",
			query:      "a=2&b=3",
			wantCode:   http.StatusOK,
			wantResult: floatptr(5),
		},
		{
			name:       "floats",
			query:      "a=1.5&b=2.25",
			wantCode:   http.StatusOK,
			wantResult: floatptr(3.75),
		},
		{
			name:       "negative operands",
			query:      "a=-4&b=10",
			wantCode:   http.StatusOK,
			wantResult: floatptr(6),
		},
		{
			name:       "zero operands",
			query:      "a=0&b=0",
			wantCode:   http.StatusOK,
			wantResult: floatptr(0),
		},
		{
			name:      "missing a",
			query:     "b=3",
			wantCode:  http.StatusBadRequest,
			wantError: `parameter "a" is missing`,
		},
		{
			name:      "missing b",
			query:     "a=3",
			wantCode:  http.StatusBadRequest,
			wantError: `parameter "b" is missing`,
		},
		{
			name:      "missing both",
			query:     "",
			wantCode:  http.StatusBadRequest,
			wantError: `parameter "a" is missing`,
		},
		{
			name:      "a not a number",
			query:     "a=abc&b=3",
			wantCode:  http.StatusBadRequest,
			wantError: `parameter "a" is not a number`,
		},
		{
			name:      "b not a number",
			query:     "a=3&b=xyz",
			wantCode:  http.StatusBadRequest,
			wantError: `parameter "b" is not a number`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://testserver/add?"+tc.query, nil)
			w := httptest.NewRecorder()

			addHandler(w, req)

			if w.Code != tc.wantCode {
				t.Fatalf("status = %d, want %d (body %q)", w.Code, tc.wantCode, w.Body.String())
			}

			if ct := w.Header().Get("Content-Type"); ct != "application/json" {
				t.Errorf("Content-Type = %q, want application/json", ct)
			}

			switch {
			case tc.wantResult != nil:
				var body resultBody
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
					t.Fatalf("decode success body %q: %v", w.Body.String(), err)
				}
				if body.Result != *tc.wantResult {
					t.Errorf("result = %v, want %v", body.Result, *tc.wantResult)
				}
			case tc.wantError != "":
				var body errorBody
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
					t.Fatalf("decode error body %q: %v", w.Body.String(), err)
				}
				if body.Error != tc.wantError {
					t.Errorf("error = %q, want %q", body.Error, tc.wantError)
				}
			}
		})
	}
}

// TestRegister verifies the public entrypoint mounts GET /add on the mux
// and that non-matching requests are rejected by standard routing.
func TestRegister(t *testing.T) {
	mux := http.NewServeMux()
	Register(mux)

	t.Run("serves a valid request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://testserver/add?a=1&b=2", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d (body %q)", w.Code, http.StatusOK, w.Body.String())
		}
	})

	t.Run("rejects a non-GET method to /add", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://testserver/add?a=1&b=2", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Fatalf("status = %d, want %d", w.Code, http.StatusMethodNotAllowed)
		}
	})

	t.Run("rejects an unknown path", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "http://testserver/not-here", nil)
		w := httptest.NewRecorder()

		mux.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("status = %d, want %d", w.Code, http.StatusNotFound)
		}
	})
}
