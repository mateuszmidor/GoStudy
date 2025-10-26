package main

import (
	"hexagons/hw"
	"hexagons/hw/infrastructure"
	"log"
	"net/http"
	"rest"
	"retry"
)

type HwAdapter struct {
	hwServicePort infrastructure.HwServicePort
}

// NewHwAdapter creates a HTTP adapter for Hw
func NewHwAdapter(hw *hw.HwRoot) HwAdapter {
	return HwAdapter{hw.GetServicePort()}
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateStationList(stationList []string) {
	endpoint := rest.MakeTunerEndpoint(rest.TunerStationList)
	retry.UntilSuccessOr5Failures("updating station list", rest.HttpPut, endpoint, stationList)
}

// Hw -> Tuner
func (adapter *HwAdapter) UpdateSubscription(subscription bool) {
	endpoint := rest.MakeTunerEndpoint(rest.TunerSubscription)
	retry.UntilSuccessOr5Failures("updating subscription", rest.HttpPut, endpoint, subscription)
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
