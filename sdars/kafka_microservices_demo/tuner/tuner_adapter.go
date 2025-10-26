package main

import (
	"bytes"
	"hexagons/tuner"
	"hexagons/tuner/infrastructure"
	"mykafka"
	"sharedkernel"

	"github.com/segmentio/kafka-go"
)

// TunerAdapter implements TunerServer generated from tuner.proto into tuner.pb.go
type TunerAdapter struct {
	tunerServicePort infrastructure.TunerServicePort // communication towards Tuner
	hwWriter         *kafka.Writer
	uiWriter         *kafka.Writer
}

// NewTunerAdapter creates a grpc adapter for Tuner
func NewTunerAdapter(tuner *tuner.TunerRoot) TunerAdapter {
	return TunerAdapter{tuner.GetServicePort(), mykafka.NewWriter(mykafka.TunerClient, mykafka.HwTopic), mykafka.NewWriter(mykafka.TunerClient, mykafka.UiTopic)}
}

// UpdateStationList makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateStationList(stations sharedkernel.StationList) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, stations) {
		mykafka.WriteMessageWithRetry5(adapter.uiWriter, mykafka.MsgUpdateStationList, buf.Bytes())
	}
}

// UpdateSubscription makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription sharedkernel.Subscription) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, subscription) {
		mykafka.WriteMessageWithRetry5(adapter.uiWriter, mykafka.MsgUpdateSubscription, buf.Bytes())
	}
}

// TuneToStation makes a call Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID sharedkernel.StationID) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, stationID) {
		mykafka.WriteMessageWithRetry5(adapter.hwWriter, mykafka.MsgTuneToStation, buf.Bytes())
	}
}

// kafkaUpdateStationList receives a call Hw -> Tuner
func (adapter *TunerAdapter) kafkaUpdateStationList(data []byte) {
	var stations []string
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &stations) {
		adapter.tunerServicePort.UpdateStationList(stations)
	}
}

// kafkaUpdateSubscription kafkaUpdateSubscription a call Hw -> Tuner
func (adapter *TunerAdapter) kafkaUpdateSubscription(data []byte) {
	var subscription bool
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &subscription) {
		adapter.tunerServicePort.UpdateSubscription(subscription)
	}
}

// kafkaTuneToStation receives a call Ui -> Tuner
func (adapter *TunerAdapter) kafkaTuneToStation(data []byte) {
	var stationID uint32
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &stationID) {
		adapter.tunerServicePort.TuneToStation(stationID)
	}
}

func (adapter *TunerAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.TunerClient, mykafka.TunerTopic)
	defer reader.Close()
	var msg kafka.Message
	var success bool

	for {
		if msg, success = mykafka.ReadMessageWithRetry5(reader); success == false {
			continue
		}

		switch string(msg.Key) {
		case mykafka.MsgUpdateStationList:
			adapter.kafkaUpdateStationList(msg.Value)
		case mykafka.MsgUpdateSubscription:
			adapter.kafkaUpdateSubscription(msg.Value)
		case mykafka.MsgTuneToStation:
			adapter.kafkaTuneToStation(msg.Value)
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for Tuner
func (adapter *TunerAdapter) RunKafkaConsumer() {
	if !mykafka.NewTopicWithRetry5(mykafka.TunerTopic, 4, 3) {
		panic("Could not register kafka topic")
	}
	adapter.readLoop()
}
