package sawmill

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	sawmillgrpc "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/grpc/gen/sawmill/v1"
	sawmillconnect "github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/grpc/gen/sawmill/v1/sawmillv1connect"
	"github.com/mateuszmidor/GoStudy/modular-monolith/sawmill/internal"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// SawmillService implements sawmillconnect.UnimplementedSawmillServiceHandler interface generated from .proto
type SawmillService struct {
	sawmillconnect.UnimplementedSawmillServiceHandler
	sawmill *internal.Sawmill
}

func RunSawmillGrpcSvc(addr string) error {
	s := internal.NewSawmill()
	s.Run()
	mux := http.NewServeMux()
	mux.Handle(sawmillconnect.NewSawmillServiceHandler(&SawmillService{sawmill: s}))
	return http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
}

// GetPlanks implements sawmillconnect.GetPlanks interface
func (s *SawmillService) GetPlanks(_ context.Context, r *connect.Request[sawmillgrpc.GetPlanksRequest]) (*connect.Response[sawmillgrpc.GetPlanksResponse], error) {
	log.Printf("server received GetPlanks request: %d", r.Msg.GetCount())
	_planks := s.sawmill.GetPlanks(int(r.Msg.Count))
	planks := make([]*sawmillgrpc.Plank, len(_planks))
	rsp := sawmillgrpc.GetPlanksResponse{Planks: planks}
	return connect.NewResponse(&rsp), nil
}

