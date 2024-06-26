// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pingpong/proto/v2alpha/ping-pong.proto

package v2alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v2alpha "github.com/mateuszmidor/GoStudy/connectgo/gen/pingpong/proto/v2alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// PingPongName is the fully-qualified name of the PingPong service.
	PingPongName = "pingpong7.PingPong"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PingPongRpcPingProcedure is the fully-qualified name of the PingPong's RpcPing RPC.
	PingPongRpcPingProcedure = "/pingpong7.PingPong/RpcPing"
)

// PingPongClient is a client for the pingpong7.PingPong service.
type PingPongClient interface {
	RpcPing(context.Context, *connect_go.Request[v2alpha.Message]) (*connect_go.Response[v2alpha.Message], error)
}

// NewPingPongClient constructs a client for the pingpong7.PingPong service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPingPongClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PingPongClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pingPongClient{
		rpcPing: connect_go.NewClient[v2alpha.Message, v2alpha.Message](
			httpClient,
			baseURL+PingPongRpcPingProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// pingPongClient implements PingPongClient.
type pingPongClient struct {
	rpcPing *connect_go.Client[v2alpha.Message, v2alpha.Message]
}

// RpcPing calls pingpong7.PingPong.RpcPing.
func (c *pingPongClient) RpcPing(ctx context.Context, req *connect_go.Request[v2alpha.Message]) (*connect_go.Response[v2alpha.Message], error) {
	return c.rpcPing.CallUnary(ctx, req)
}

// PingPongHandler is an implementation of the pingpong7.PingPong service.
type PingPongHandler interface {
	RpcPing(context.Context, *connect_go.Request[v2alpha.Message]) (*connect_go.Response[v2alpha.Message], error)
}

// NewPingPongHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPingPongHandler(svc PingPongHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	pingPongRpcPingHandler := connect_go.NewUnaryHandler(
		PingPongRpcPingProcedure,
		svc.RpcPing,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/pingpong7.PingPong/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PingPongRpcPingProcedure:
			pingPongRpcPingHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPingPongHandler returns CodeUnimplemented from all methods.
type UnimplementedPingPongHandler struct{}

func (UnimplementedPingPongHandler) RpcPing(context.Context, *connect_go.Request[v2alpha.Message]) (*connect_go.Response[v2alpha.Message], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pingpong7.PingPong.RpcPing is not implemented"))
}
