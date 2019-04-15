package vault

import (
	context "context"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(
		ctx, endpoints.HashEndpoint, decodeHashRequest, encodeResponse))
	)
	m.Handle("/validate", httptransport.NewServer(c
		tx, endpoints.ValidateEndpoint, decodeValidateRequest, encodeResponse))
	return m
}
