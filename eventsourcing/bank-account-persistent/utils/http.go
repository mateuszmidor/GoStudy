package utils

import (
	"fmt"
	"log/slog"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.statusCode = code
	sw.ResponseWriter.WriteHeader(code)
}

func NewRequestLogger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}
		mux.ServeHTTP(sw, r)
		paddedMethod := fmt.Sprintf("%4s", r.Method)
		s := fmt.Sprintf("[%d] %s %s", sw.statusCode, paddedMethod, r.URL.Path)
		slog.Info(s)
	})
}
