package clients

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sailworksgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/api/gen/sailworks/v1"
	sailworksconnect "github.com/mateuszmidor/GoStudy/modular-monolith/api/gen/sailworks/v1/sailworksv1connect"
)

// SailworksGRPC implements the Sailworks interface and exposes SailWorksSvc over Grpc
type SailworksGRPC struct {
	client sailworksconnect.SailworksServiceClient
}

func NewSailworksGrpc(addr string) *SailworksGRPC {
	log.Println("NewSailworksGRPC client:", addr)
	client := sailworksconnect.NewSailworksServiceClient(http.DefaultClient, "http://"+addr)
	return &SailworksGRPC{client: client}
}

func (sg *SailworksGRPC) GetSails(count int) ([]Sail, error) {
	msg := sailworksgrpc.GetSailsRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetSails(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SailworksGrpc GetSails failed")
	}
	return make([]Sail, len(rsp.Msg.Sails)), nil
}

func (sl *SailworksGRPC) Run() {
	// nothing to do as Sailworks should be running as a separate process
}
