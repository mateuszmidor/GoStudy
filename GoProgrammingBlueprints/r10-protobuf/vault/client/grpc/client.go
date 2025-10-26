package grpc

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/r10-protobuf/vault"
	"github.com/mateuszmidor/GoStudy/GoProgrammingBlueprints/r10-protobuf/vault/pb"
	"google.golang.org/grpc"
)

// New creates Vault Service over GRPC
func New(conn *grpc.ClientConn) vault.Service {
	var hashEndpoint = grpctransport.NewClient(
		conn, "pb.Vault", "Hash",
		vault.EncodeGRPCHashRequest,
		vault.DecodeGRPCHashResponse,
		pb.HashResponse{},
	).Endpoint()

	var validateEndpoint = grpctransport.NewClient(
		conn, "pb.Vault", "Validate",
		vault.EncodeGRPCValidateRequest,
		vault.DecodeGRPCValidateResponse,
		pb.ValidateResponse{},
	).Endpoint()

	return vault.Endpoints{
		HashEndpoint:     hashEndpoint,
		ValidateEndpoint: validateEndpoint,
	}
}
