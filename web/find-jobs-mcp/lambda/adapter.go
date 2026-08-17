package main

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mark3labs/mcp-go/server"
)

// mcpAdapter bridges Lambda Function URL events into the MCP streamable HTTP
// handler and converts the response back into a Lambda response.
type mcpAdapter struct {
	handler http.Handler
}

// newMCPAdapter wraps the MCP server in a stateless, non-streaming
// streamable HTTP handler that fits Lambda's request/response model:
//   - stateless: no session survives across invocations, every request is a fresh session
//   - no streaming: GET (SSE notification streams) replies 405 instead of hanging
func newMCPAdapter(mcpServer *server.MCPServer) *mcpAdapter {
	return &mcpAdapter{
		handler: server.NewStreamableHTTPServer(
			mcpServer,
			server.WithStateLess(true),
			server.WithDisableStreaming(true),
		),
	}
}

// handle is the Lambda entry point. It converts a Function URL event into an
// http.Request, serves it through the MCP handler and maps the result back.
func (a *mcpAdapter) handle(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	httpReq, err := toHTTPRequest(ctx, req)
	if err != nil {
		var reqErr *requestError
		if errors.As(err, &reqErr) {
			return errorResponse(reqErr.status, reqErr.Error()), nil
		}
		return errorResponse(http.StatusBadRequest, "bad request: "+err.Error()), nil
	}

	rec := httptest.NewRecorder()
	a.handler.ServeHTTP(rec, httpReq)

	return events.LambdaFunctionURLResponse{
		StatusCode: rec.Code,
		Headers:    toHeaderMap(rec.Header()),
		Body:       rec.Body.String(),
	}, nil
}

// requestError carries the HTTP status to return for a malformed request.
type requestError struct {
	status int
	msg    string
}

func (e *requestError) Error() string { return e.msg }

// toHTTPRequest converts a Function URL event into an http.Request for the
// MCP endpoint /mcp.
func toHTTPRequest(ctx context.Context, req events.LambdaFunctionURLRequest) (*http.Request, error) {
	if req.RawPath != "/mcp" {
		return nil, &requestError{status: http.StatusNotFound, msg: "not found: expected /mcp"}
	}

	body := req.Body
	if req.IsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			return nil, &requestError{status: http.StatusBadRequest, msg: "invalid base64 body: " + err.Error()}
		}
		body = string(decoded)
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.RequestContext.HTTP.Method, req.RawPath, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}
	return httpReq, nil
}

// toHeaderMap flattens an http.Header into the string map expected by Lambda.
func toHeaderMap(h http.Header) map[string]string {
	out := make(map[string]string, len(h))
	for k, v := range h {
		out[k] = strings.Join(v, ", ")
	}
	return out
}

func errorResponse(status int, body string) events.LambdaFunctionURLResponse {
	return events.LambdaFunctionURLResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       body,
	}
}
