package main

import (
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"net/http"
)

type HwAdapter struct {
	hwServicePort infrastructure.ServicePort
}

// NewHwAdapter creates a HTTP adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort()}
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	httpPut(MakeTunerEndpoint(TunerStationList), stationList)
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	httpPut(MakeTunerEndpoint(TunerSubscription), subscription)
}

func (adapter *HwAdapter) handleCurrentStation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var currentStationID uint32
		if err := decodeBody(r, &currentStationID); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.hwServicePort.TuneToStation(currentStationID)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for Hw
func (adapter *HwAdapter) RunHTTPServer() {
	addr := HwAddr
	mux := http.NewServeMux()
	mux.HandleFunc(HwCurrentStation, adapter.handleCurrentStation)
	log.Println("Starting HwAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping HwAdapterr...")
}
