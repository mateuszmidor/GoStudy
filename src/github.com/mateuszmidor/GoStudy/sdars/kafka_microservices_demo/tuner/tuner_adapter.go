package main

import (
	"bytes"
	"fmt"
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"log"
	"mykafka"

	"github.com/segmentio/kafka-go"
)

// TunerAdapter implements TunerServer generated from tuner.proto into tuner.pb.go
type TunerAdapter struct {
	tunerServicePort infrastructure.ServicePort // communication towards Tuner
	hwWriter         *kafka.Writer
	uiWriter         *kafka.Writer
}

// NewTunerAdapter creates a grpc adapter for Tuner
func NewTunerAdapter(tuner *tuner.TunerRoot) TunerAdapter {
	return TunerAdapter{tuner.GetServicePort(), mykafka.NewWriter(mykafka.TunerClient, mykafka.HwTopic), mykafka.NewWriter(mykafka.TunerClient, mykafka.UiTopic)}
}

// UpdateStationList makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateStationList(stations domain.StationList) {
	buf := bytes.NewBuffer([]byte{})
	mykafka.EncodeBody(buf, stations)
	mykafka.Write(adapter.uiWriter, mykafka.MsgUpdateStationList, buf.Bytes())
}

// UpdateSubscription makes a call Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription domain.Subscription) {
	buf := bytes.NewBuffer([]byte{})
	mykafka.EncodeBody(buf, subscription)
	mykafka.Write(adapter.uiWriter, mykafka.MsgUpdateSubscription, buf.Bytes())
}

// TuneToStation makes a call Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID domain.StationId) {
	buf := bytes.NewBuffer([]byte{})
	mykafka.EncodeBody(buf, stationID)
	mykafka.Write(adapter.hwWriter, mykafka.MsgTuneToStation, buf.Bytes())
}

// kafkaUpdateStationList receives a call Hw -> Tuner
func (adapter *TunerAdapter) kafkaUpdateStationList(data []byte) {
	var stations []string
	mykafka.DecodeBody(bytes.NewReader(data), &stations)
	adapter.tunerServicePort.StationListUpdated(stations)
}

// kafkaUpdateSubscription kafkaUpdateSubscription a call Hw -> Tuner
func (adapter *TunerAdapter) kafkaUpdateSubscription(data []byte) {
	var subscription bool
	mykafka.DecodeBody(bytes.NewReader(data), &subscription)
	adapter.tunerServicePort.SubscriptionUpdated(subscription)
}

// kafkaTuneToStation receives a call Ui -> Tuner
func (adapter *TunerAdapter) kafkaTuneToStation(data []byte) {
	var stationId uint32
	mykafka.DecodeBody(bytes.NewReader(data), &stationId)
	adapter.tunerServicePort.TuneToStation(stationId)
}

func (adapter *TunerAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.TunerClient, mykafka.TunerTopic)
	defer reader.Close()

	for {
		m, err := mykafka.Read(reader)
		if err != nil {
			log.Printf("error while receiving message: %s\n", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s: %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		switch string(m.Key) {
		case mykafka.MsgUpdateStationList:
			adapter.kafkaUpdateStationList(m.Value)
		case mykafka.MsgUpdateSubscription:
			adapter.kafkaUpdateSubscription(m.Value)
		case mykafka.MsgTuneToStation:
			adapter.kafkaTuneToStation(m.Value)
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for Tuner
func (adapter *TunerAdapter) RunKafkaConsumer() {
	mykafka.NewTopic(mykafka.TunerTopic, 4, 2)
	adapter.readLoop()
}
