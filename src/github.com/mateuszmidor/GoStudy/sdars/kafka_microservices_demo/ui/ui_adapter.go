package main

import (
	"bytes"
	"fmt"
	"hexagons/ui"
	"hexagons/ui/infrastructure"
	"log"
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
	mykafka.EncodeBody(buf, stationID)
	mykafka.Write(adapter.tunerWriter, mykafka.MsgTuneToStation, buf.Bytes())
}

// kafkaUpdateStationList forwards a call Tuner -> Ui
func (adapter *UIAdapter) kafkaUpdateStationList(data []byte) {
	var stations []string
	mykafka.DecodeBody(bytes.NewReader(data), &stations)
	adapter.uiServicePort.UpdateStationList(stations)
}

// kafkaUpdateSubscription forwards a call Tuner -> Ui
func (adapter *UIAdapter) kafkaUpdateSubscription(data []byte) {
	var subscription bool
	mykafka.DecodeBody(bytes.NewReader(data), &subscription)
	adapter.uiServicePort.UpdateSubscription(subscription)
}

func (adapter *UIAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.UiClient, mykafka.UiTopic)
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
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for UI
func (adapter *UIAdapter) RunKafkaConsumer() {
	mykafka.NewTopic(mykafka.UiTopic, 4, 2)
	adapter.readLoop()
}
