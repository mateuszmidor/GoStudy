// Package api defines the common JSON wire format shared by all
// slice endpoints of the go-arch-lint demo.
package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ResultSuccess is the response body returned by endpoints on success.
type ResultSuccess struct {
	Result float64 `json:"result"`
}

// ResultError is the response body returned by endpoints on failure.
type ResultError struct {
	Error string `json:"error"`
}

// ReturnSuccess writes a 200 JSON response containing result to w.
func ReturnSuccess(w http.ResponseWriter, result float64) {
	writeJSON(w, http.StatusOK, ResultSuccess{Result: result})
}

// ReturnError writes a JSON response containing message
// with the given HTTP status code to w.
func ReturnError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ResultError{Error: message})
}

// writeJSON sets the JSON content type, writes the status code
// and encodes v into the response.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	// The error is ignored: headers are already written at this point,
	// so there is nothing meaningful left to do on failure.
	_ = json.NewEncoder(w).Encode(v)
}

// ParseParam reads the required query parameter name
// and parses it as a float64.
func ParseParam(q url.Values, name string) (float64, error) {
	raw := q.Get(name)
	if raw == "" {
		return 0, &ParamErr{Reason: "parameter \"" + name + "\" is missing"}
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, &ParamErr{Reason: "parameter \"" + name + "\" is not a number"}
	}
	return v, nil
}

// ParamErr is a minimal error type for missing or malformed parameters.
type ParamErr struct{ Reason string }

func (e *ParamErr) Error() string { return e.Reason }
