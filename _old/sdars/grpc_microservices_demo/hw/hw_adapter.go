package main

import (
	"context"
	"errors"
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"retry"
	"rpc"
)

// HwAdapter implements HwServer generated from hw.proto into hw.pb.go
type HwAdapter struct {
	hwServicePort infrastructure.HwServicePort // communication towards Hw
	tunerClient   rpc.TunerClient              // communication towards Tuner
}

// NewHwAdapter creates a grpc adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort(), nil}
}

// UpdateStationList makes a call Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	f := func() error {
		if adapter.tunerClient == nil {
			return errors.New("tuner not available")
		}

		rq := &rpc.TunerUpdateStationListRequest{}
		rq.Stations = stationList
		_, err := adapter.tunerClient.RpcUpdateStationList(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("updating station list", f)
}

// UpdateSubscription makes a call Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	f := func() error {
		if adapter.tunerClient == nil {
			return errors.New("tuner not available")
		}

		rq := &rpc.TunerUpdateSubscriptionRequest{}
		rq.Active = subscription
		_, err := adapter.tunerClient.RpcUpdateSubscription(context.Background(), rq)
		return err
	}
	retry.UntilSuccessOr5Failures("updating subscription", f)
}

// RpcTuneToStation receives a call Tuner -> Hw
func (adapter *HwAdapter) RpcTuneToStation(ctx context.Context, rq *rpc.HwTuneToStationRequest) (*rpc.Empty, error) {
	adapter.hwServicePort.TuneToStation(rq.StationID)
	return &rpc.Empty{}, nil
}

// RunGrpcServer starts a server that receives calls from Tuner
func (adapter *HwAdapter) RunGrpcServer() {
	// create connection towards Tuner
	tunerConn := rpc.ConnectGrpcClient(rpc.TunerAddr)
	defer tunerConn.Close()
	adapter.tunerClient = rpc.NewTunerClient(tunerConn)

	// config & run server
	lis, grpcServer := rpc.MakeGrpcServer(rpc.HwAddr)
	rpc.RegisterHwServer(grpcServer, adapter)
	grpcServer.Serve(lis)
	log.Println("hw listen done")
}
