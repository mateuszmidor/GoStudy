package main

import (
	"bytes"
	"hexagons/ui"
	"hexagons/ui/infrastructure"
	"mykafka"

	"github.com/segmentio/kafka-go"
)

// UIAdapter implements tuner output ports towards ui, and ui output ports towards tuner
type UIAdapter struct {
	uiServicePort infrastructure.ServicePort
	tunerWriter   *kafka.Writer
}

// NewUIAdapter creates a HTTP adapter for UI
func NewUIAdapter(ui *ui.UiRoot) UIAdapter {
	return UIAdapter{ui.GetServicePort(), mykafka.NewWriter(mykafka.UiClient, mykafka.TunerTopic)}
}

// TuneToStation forwards command UI -> Tuner
func (adapter *UIAdapter) TuneToStation(stationID uint32) {
	buf := bytes.NewBuffer([]byte{})
	if mykafka.EncodeMessageOrLog(buf, stationID) {
		mykafka.WriteMessageWithRetry5(adapter.tunerWriter, mykafka.MsgTuneToStation, buf.Bytes())
	}
}

// kafkaUpdateStationList forwards a call Tuner -> Ui
func (adapter *UIAdapter) kafkaUpdateStationList(data []byte) {
	var stations []string
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &stations) {
		adapter.uiServicePort.UpdateStationList(stations)
	}
}

// kafkaUpdateSubscription forwards a call Tuner -> Ui
func (adapter *UIAdapter) kafkaUpdateSubscription(data []byte) {
	var subscription bool
	if mykafka.DecodeMessageOrLog(bytes.NewReader(data), &subscription) {
		adapter.uiServicePort.UpdateSubscription(subscription)
	}
}

func (adapter *UIAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.UiClient, mykafka.UiTopic)
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
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for UI
func (adapter *UIAdapter) RunKafkaConsumer() {
	if !mykafka.NewTopicWithRetry5(mykafka.UiTopic, 4, 3) {
		panic("Could not register kafka topic")
	}
	adapter.readLoop()
}
