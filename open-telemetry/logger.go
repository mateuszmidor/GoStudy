package main

import (
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

func newLogger(scopeName string) *slog.Logger {
	return slog.New(otelslog.NewHandler(scopeName))
}
