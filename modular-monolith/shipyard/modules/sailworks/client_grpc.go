package sailworks

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sailworksgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks/grpc/gen/sailworks/v1"
	sailworksconnect "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks/grpc/gen/sailworks/v1/sailworksv1connect"
)

// APIGrpc implements the sailworks module API as GRPC client.
type APIGrpc struct {
	client sailworksconnect.SailworksServiceClient
}

func NewSailworksGRPC(addr string) *APIGrpc {
	log.Println("NewSailworksGrpc client:", addr)
	client := sailworksconnect.NewSailworksServiceClient(http.DefaultClient, "http://"+addr)
	return &APIGrpc{client: client}
}

func (sg *APIGrpc) GetSails(count int) ([]Sail, error) {
	msg := sailworksgrpc.GetSailsRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetSails(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SailworksGrpc GetSails failed")
	}
	return make([]Sail, len(rsp.Msg.Sails)), nil
}

func (sl *APIGrpc) Run() {
	// nothing to do as Sailworks should be running as a separate process
}
