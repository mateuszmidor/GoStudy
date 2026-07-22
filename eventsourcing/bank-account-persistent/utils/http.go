package utils

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

func (sw *statusWriter) Write(b []byte) (int, error) {
	sw.body.Write(b)
	return sw.ResponseWriter.Write(b)
}

// NewRequestLogger wraps "mux" to intercept and log incoming http requests and their responses
func NewRequestLogger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		mux.ServeHTTP(sw, r)
		paddedMethodName := fmt.Sprintf("%4s", r.Method)
		reqStatus := fmt.Sprintf("[%d] %s %s", sw.statusCode, paddedMethodName, r.URL.Path)
		if sw.statusCode >= 400 {
			slog.Error(reqStatus, "body", sw.body.String())
		} else {
			slog.Info(reqStatus)
		}
	})
}
