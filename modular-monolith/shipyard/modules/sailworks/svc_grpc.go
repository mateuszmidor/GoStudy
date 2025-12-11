package sailworks

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	sailworksgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/api/gen/sailworks/v1"
	sailworksconnect "github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/api/gen/sailworks/v1/sailworksv1connect"
	"github.com/mateuszmidor/GoStudy/modular-monolith/shipyard/modules/sailworks/internal"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// SailWorksSvc implements sailworksconnect.UnimplementedSailworksServiceHandler interface generated from .proto
type SailWorksSvc struct {
	sailworksconnect.UnimplementedSailworksServiceHandler
	_sailworks *internal.Sailworks
}

func RunSailworksGrpcSvc(addr string) error {
	s := internal.NewSailworks()
	s.Run()
	mux := http.NewServeMux()
	mux.Handle(sailworksconnect.NewSailworksServiceHandler(&SailWorksSvc{_sailworks: s}))
	return http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
}

// GetSails implements sailworksconnect.GetSails interface
func (s *SailWorksSvc) GetSails(_ context.Context, r *connect.Request[sailworksgrpc.GetSailsRequest]) (*connect.Response[sailworksgrpc.GetSailsResponse], error) {
	log.Printf("server received GetSails request: %d", r.Msg.GetCount())
	_sails := s._sailworks.GetSails(int(r.Msg.Count))
	sails := make([]*sailworksgrpc.Sail, len(_sails))
	rsp := sailworksgrpc.GetSailsResponse{Sails: sails}
	return connect.NewResponse(&rsp), nil
}
