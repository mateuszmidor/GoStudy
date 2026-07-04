package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	sawmillgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/api/grpc/gen/sawmill/v1"
	sawmillconnect "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/api/grpc/gen/sawmill/v1/sawmillv1connect"
	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/cmd/sawmill/internal"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type SawmillService struct {
	sawmillconnect.UnimplementedSawmillServiceHandler
	sawmill *internal.Sawmill
}

func RunGrpcService(addr string) error {
	sawmill := internal.NewSawmill()
	sawmill.Run()
	mux := http.NewServeMux()
	mux.Handle(sawmillconnect.NewSawmillServiceHandler(&SawmillService{sawmill: sawmill}))
	return http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
}

func (s *SawmillService) GetBeams(_ context.Context, r *connect.Request[sawmillgrpc.GetBeamsRequest]) (*connect.Response[sawmillgrpc.GetBeamsResponse], error) {
	log.Printf("SawmillService received GetBeams request: %d", r.Msg.GetCount())
	_beams := s.sawmill.GetBeams(int(r.Msg.Count))
	beams := make([]*sawmillgrpc.Beam, len(_beams))
	rsp := sawmillgrpc.GetBeamsResponse{Beams: beams}
	return connect.NewResponse(&rsp), nil
}
