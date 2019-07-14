package main

import (
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"net/http"
	"rest"
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
	rest.HttpPut(rest.MakeTunerEndpoint(rest.TunerStationList), stationList)
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	rest.HttpPut(rest.MakeTunerEndpoint(rest.TunerSubscription), subscription)
}

func (adapter *HwAdapter) handleCurrentStation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var currentStationID uint32
		if err := rest.DecodeBody(r, &currentStationID); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.hwServicePort.TuneToStation(currentStationID)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for Hw
func (adapter *HwAdapter) RunHTTPServer() {
	addr := rest.HwAddr
	mux := http.NewServeMux()
	mux.HandleFunc(rest.HwCurrentStation, adapter.handleCurrentStation)
	log.Println("Starting HwAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping HwAdapterr...")
}
