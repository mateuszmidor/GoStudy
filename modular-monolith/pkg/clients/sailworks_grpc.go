package clients

import (
	"context"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/bufbuild/connect-go"
	sailworksgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/api/gen/sailworks/v1"
	sailworksconnect "github.com/mateuszmidor/GoStudy/modular-monolith/api/gen/sailworks/v1/sailworksv1connect"
	"github.com/mateuszmidor/GoStudy/modular-monolith/internal/modules/sailworks"
)

// SailworksGrpc implements the Sailworks interface and exposes SailWorksSvc over Grpc
type SailworksGrpc struct {
	client sailworksconnect.SailworksServiceClient
}

func NewSailworksGrpc(addr string) *SailworksGrpc {
	log.Println("NewSailworksGrpc client:", addr)
	if !pingGrpcAddr(addr) {
		log.Fatal("no service running at: ", addr)
	}
	client := sailworksconnect.NewSailworksServiceClient(http.DefaultClient, "http://"+addr)
	return &SailworksGrpc{client: client}
}

func (sg *SailworksGrpc) GetSails(count int) ([]sailworks.Sail, error) {
	msg := sailworksgrpc.GetSailsRequest{Count: int32(count)}
	req := connect.NewRequest(&msg)
	rsp, err := sg.client.GetSails(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "SailworksGrpc GetSails failed")
	}
	return make([]sailworks.Sail, len(rsp.Msg.Sails)), nil
}

func (sl *SailworksGrpc) Run() {
	// nothing to do as Sailworks should be running as a separate process
}

func pingGrpcAddr(addr string) bool {
	_, err := http.Get("http://" + addr)
	if err != nil {
		log.Println(err) // if no svc is running at "addr", the err will be: connection refused
		return false
	}

	return true
}
