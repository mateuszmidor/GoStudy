package sawmill

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sawmillgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/grpc/gen/sawmill/v1"
	sawmillconnect "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/grpc/gen/sawmill/v1/sawmillv1connect"
)

// APIGrpc implements the sawmill module API as GRPC client.
type APIGrpc struct {
	client sawmillconnect.SawmillServiceClient
}

func NewSawmillGRPC(addr string) *APIGrpc {
	log.Println("NewSawmillGrpc client:", addr)
	client := sawmillconnect.NewSawmillServiceClient(http.DefaultClient, "http://"+addr)
	return &APIGrpc{client: client}
}

func (sg *APIGrpc) GetPlanks(count int) ([]Plank, error) {
	msg := sawmillgrpc.GetPlanksRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetPlanks(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SawmillGrpc GetPlanks failed")
	}
	return make([]Plank, len(rsp.Msg.Planks)), nil
}

func (sg *APIGrpc) Run() {
	// nothing to do as Sawmill should be running as a separate process
}

