package sawmill

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sawmillgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/api/grpc/gen/sawmill/v1"
	sawmillconnect "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/api/grpc/gen/sawmill/v1/sawmillv1connect"
)

// GrpcClient implements the sawmill module API as GRPC client.
type GrpcClient struct {
	client sawmillconnect.SawmillServiceClient
}

func NewGrpcClient(addr string) *GrpcClient {
	log.Println("NewGRPCAPI sawmill client:", addr)
	client := sawmillconnect.NewSawmillServiceClient(http.DefaultClient, "http://"+addr)
	return &GrpcClient{client: client}
}

func (sg *GrpcClient) GetBeams(count int) ([]Beam, error) {
	msg := sawmillgrpc.GetBeamsRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetBeams(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SawmillGrpc GetBeams failed")
	}
	return make([]Beam, len(rsp.Msg.Beams)), nil
}
