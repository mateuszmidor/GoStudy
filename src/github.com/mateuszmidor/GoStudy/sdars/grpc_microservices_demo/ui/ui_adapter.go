package main

import (
	"context"
	"errors"
	"hexagons/ui"
	"hexagons/ui/infrastructure"
	"log"
	"retry"
	"rpc"
)

// UIAdapter implements UiServer generated from ui.proto into ui.pb.go
type UIAdapter struct {
	uiServicePort infrastructure.UiServicePort // communication towards Ui
	tunerClient   rpc.TunerClient              // communication towards Tuner
}

// NewUIAdapter creates a grpc adapter for Ui
func NewUIAdapter(ui *ui.UiRoot) UIAdapter {
	return UIAdapter{ui.GetServicePort(), nil}
}

// TuneToStation makes a call Ui -> Tuner
func (adapter *UIAdapter) TuneToStation(stationID uint32) {
	f := func() error {
		if adapter.tunerClient == nil {
			return errors.New("tuner not available")
		}

		rq := &rpc.TunerTuneToStationRequest{}
		rq.StationID = stationID
		_, err := adapter.tunerClient.RpcTuneToStation(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("tuning to station", f)
}

// RpcUpdateStationList receives a call Tuner -> Ui
func (adapter *UIAdapter) RpcUpdateStationList(ctx context.Context, rq *rpc.UIUpdateStationListRequest) (*rpc.Empty, error) {
	adapter.uiServicePort.UpdateStationList(rq.Stations)
	return &rpc.Empty{}, nil
}

// RpcUpdateSubscription receives a call Tuner -> Ui
func (adapter *UIAdapter) RpcUpdateSubscription(ctx context.Context, rq *rpc.UIUpdateSubscriptionRequest) (*rpc.Empty, error) {
	adapter.uiServicePort.UpdateSubscription(rq.Active)
	return &rpc.Empty{}, nil
}

// RunGrpcServer starts a server that handles commands from Tuner
func (adapter *UIAdapter) RunGrpcServer() {
	// create connection towards Tuner
	tunerConn := rpc.ConnectGrpcClient(rpc.TunerAddr)
	defer tunerConn.Close()
	adapter.tunerClient = rpc.NewTunerClient(tunerConn)

	// config & run server
	lis, grpcServer := rpc.MakeGrpcServer(rpc.UIAddr)
	rpc.RegisterUIServer(grpcServer, adapter)
	grpcServer.Serve(lis)
	log.Println("ui listen done")
}
