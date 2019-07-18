package main

import (
	"context"
	"fmt"
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"rpc"
)

// HwAdapter implements HwServer generated from hw.proto into hw.pb.go
type HwAdapter struct {
	hwServicePort infrastructure.ServicePort // communication towards Hw
	tunerClient   rpc.TunerClient            // communication towards Tuner
}

// NewHwAdapter creates a grpc adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort(), nil}
}

// UpdateStationList makes a call Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	fmt.Println("HwAdapter.UpdateStationList -> Tuner")
	if adapter.tunerClient == nil {
		fmt.Println("tuner not available")
		return
	}

	rq := &rpc.TunerUpdateStationListRequest{}
	rq.Stations = stationList
	_, err := adapter.tunerClient.RpcUpdateStationList(context.Background(), rq)

	rpc.LogCallResult(err)
}

// UpdateSubscription makes a call Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	fmt.Println("HwAdapter.UpdateSubscription -> Tuner")
	if adapter.tunerClient == nil {
		fmt.Println("tuner not available")
		return
	}

	rq := &rpc.TunerUpdateSubscriptionRequest{}
	rq.Active = subscription
	_, err := adapter.tunerClient.RpcUpdateSubscription(context.Background(), rq)

	rpc.LogCallResult(err)
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
