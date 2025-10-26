package main

import (
	"context"
	"errors"
	"hexagons/tuner"
	"hexagons/tuner/infrastructure"
	"log"
	"retry"
	"rpc"
	"sharedkernel"
)

// TunerAdapter implements TunerServer generated from tuner.proto into tuner.pb.go
type TunerAdapter struct {
	tunerServicePort infrastructure.TunerServicePort // communication towards Tuner
	hwClient         rpc.HwClient                    // communication towards Hw
	uiClient         rpc.UIClient                    // communication towards Ui
}

// NewTunerAdapter creates a grpc adapter for Tuner
func NewTunerAdapter(tuner *tuner.TunerRoot) TunerAdapter {
	return TunerAdapter{tuner.GetServicePort(), nil, nil}
}

// UpdateStationList makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateStationList(stations sharedkernel.StationList) {
	f := func() error {
		if adapter.uiClient == nil {
			return errors.New("ui not available")
		}

		rq := &rpc.UIUpdateStationListRequest{}
		rq.Stations = stations
		_, err := adapter.uiClient.RpcUpdateStationList(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("updating station list", f)
}

// UpdateSubscription makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription sharedkernel.Subscription) {
	f := func() error {
		if adapter.uiClient == nil {
			return errors.New("ui not available")
		}

		rq := &rpc.UIUpdateSubscriptionRequest{}
		rq.Active = subscription
		_, err := adapter.uiClient.RpcUpdateSubscription(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("updating subscription", f)
}

// TuneToStation makes a call Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID sharedkernel.StationID) {
	f := func() error {
		if adapter.hwClient == nil {
			return errors.New("hw not available")
		}

		rq := &rpc.HwTuneToStationRequest{}
		rq.StationID = stationID
		_, err := adapter.hwClient.RpcTuneToStation(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("tuning to station", f)
}

// RpcUpdateStationList receives a call Hw -> Tuner
func (adapter *TunerAdapter) RpcUpdateStationList(_ context.Context, rq *rpc.TunerUpdateStationListRequest) (*rpc.Empty, error) {
	adapter.tunerServicePort.UpdateStationList(rq.GetStations())
	return &rpc.Empty{}, nil
}

// RpcUpdateSubscription receives a call Hw -> Tuner
func (adapter *TunerAdapter) RpcUpdateSubscription(_ context.Context, rq *rpc.TunerUpdateSubscriptionRequest) (*rpc.Empty, error) {
	adapter.tunerServicePort.UpdateSubscription(rq.GetActive())
	return &rpc.Empty{}, nil
}

// RpcUpdateSubscription receives a call Ui -> Tuner
func (adapter *TunerAdapter) RpcTuneToStation(_ context.Context, rq *rpc.TunerTuneToStationRequest) (*rpc.Empty, error) {
	adapter.tunerServicePort.TuneToStation(rq.GetStationID())
	return &rpc.Empty{}, nil
}

// RunGrpcServer starts a server that handles commands for Tuner
func (adapter *TunerAdapter) RunGrpcServer() {
	// create connection towards Hw
	hwConn := rpc.ConnectGrpcClient(rpc.HwAddr)
	defer hwConn.Close()
	adapter.hwClient = rpc.NewHwClient(hwConn)

	// create connection towards Ui
	uiConn := rpc.ConnectGrpcClient(rpc.UIAddr)
	defer uiConn.Close()
	adapter.uiClient = rpc.NewUIClient(uiConn)

	// config & run server
	lis, grpcServer := rpc.MakeGrpcServer(rpc.TunerAddr)
	rpc.RegisterTunerServer(grpcServer, adapter)
	grpcServer.Serve(lis)
	log.Println("tuner listen done")
}
