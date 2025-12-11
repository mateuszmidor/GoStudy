package sailworks

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sailworksgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/api/gen/sailworks/v1"
	sailworksconnect "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/api/gen/sailworks/v1/sailworksv1connect"
)

// SailworksGrpc implements the Sailworks interface and exposes SailWorksSvc over Grpc
type SailworksGrpc struct {
	client sailworksconnect.SailworksServiceClient
}

func NewGrpcClient(addr string) *SailworksGrpc {
	log.Println("NewSailworksGrpc client:", addr)
	client := sailworksconnect.NewSailworksServiceClient(http.DefaultClient, "http://"+addr)
	return &SailworksGrpc{client: client}
}

func (sg *SailworksGrpc) GetSails(count int) ([]Sail, error) {
	msg := sailworksgrpc.GetSailsRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetSails(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SailworksGrpc GetSails failed")
	}
	return make([]Sail, len(rsp.Msg.Sails)), nil
}

func (sl *SailworksGrpc) Run() {
	// nothing to do as Sailworks should be running as a separate process
}
