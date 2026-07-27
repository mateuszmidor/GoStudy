package main

import (
	"bytes"
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"mykafka"

	"github.com/segmentio/kafka-go"
)

type HwAdapter struct {
	hwServicePort infrastructure.HwServicePort
	tunerWriter   *kafka.Writer
}

// NewHwAdapter creates a Kafka adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort(), mykafka.NewWriter(mykafka.HwClient, mykafka.TunerTopic)}
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, stationList) {
		mykafka.WriteMessageWithRetry5(adapter.tunerWriter, mykafka.MsgUpdateStationList, buf.Bytes())
	}
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, subscription) {
		mykafka.WriteMessageWithRetry5(adapter.tunerWriter, mykafka.MsgUpdateSubscription, buf.Bytes())
	}
}

// Tuner -> Hw
func (adapter *HwAdapter) kafkaTuneToStation(data []byte) {
	var stationID uint32
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &stationID) {
		adapter.hwServicePort.TuneToStation(stationID)
	}
}

func (adapter *HwAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.HwClient, mykafka.HwTopic)
	defer reader.Close()
	var msg kafka.Message
	var success bool

	for {
		if msg, success = mykafka.ReadMessageWithRetry5(reader); success == false {
			continue
		}
		switch string(msg.Key) {
		case mykafka.MsgTuneToStation:
			adapter.kafkaTuneToStation(msg.Value)
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for Hw
func (adapter *HwAdapter) RunKafkaConsumer() {
	if !mykafka.NewTopicWithRetry5(mykafka.HwTopic, 4, 3) {
		panic("Could not register kafka topic")
	}
	adapter.readLoop()
}
