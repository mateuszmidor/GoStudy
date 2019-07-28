package main

import (
	"bytes"
	"fmt"
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"mykafka"

	"github.com/segmentio/kafka-go"
)

type HwAdapter struct {
	hwServicePort infrastructure.ServicePort
	tunerWriter   *kafka.Writer
}

// NewHwAdapter creates a Kafka adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort(), mykafka.NewWriter(mykafka.HwClient, mykafka.TunerTopic)}
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	buf := bytes.NewBuffer([]byte{})
	mykafka.EncodeBody(buf, stationList)
	mykafka.Write(adapter.tunerWriter, mykafka.MsgUpdateStationList, buf.Bytes())
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	buf := bytes.NewBuffer([]byte{})
	mykafka.EncodeBody(buf, subscription)
	mykafka.Write(adapter.tunerWriter, mykafka.MsgUpdateSubscription, buf.Bytes())
}

// Tuner -> Hw
func (adapter *HwAdapter) kafkaTuneToStation(data []byte) {
	var stationID uint32
	mykafka.DecodeBody(bytes.NewReader(data), &stationID)
	adapter.hwServicePort.TuneToStation(stationID)
}

func (adapter *HwAdapter) readLoop() {
	reader := mykafka.NewReader(mykafka.HwClient, mykafka.HwTopic)
	defer reader.Close()

	for {
		m, err := mykafka.Read(reader)
		if err != nil {
			log.Printf("error while receiving message: %s\n", err.Error())
			continue
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s: %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		switch string(m.Key) {
		case mykafka.MsgTuneToStation:
			adapter.kafkaTuneToStation(m.Value)
		}
	}
}

// RunKafkaConsumer starts a consumer that fetches commands for Hw
func (adapter *HwAdapter) RunKafkaConsumer() {
	mykafka.NewTopic(mykafka.HwTopic, 4, 2)
	adapter.readLoop()
}
