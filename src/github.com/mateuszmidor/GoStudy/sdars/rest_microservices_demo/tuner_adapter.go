package main

import (
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"log"
	"net/http"
)

type TunerAdapter struct {
	tunerServicePort infrastructure.ServicePort
}

// NewTunerAdapter creates a HTTP adapter for Tuner hexagon
func NewTunerAdapter(tuner *tuner.TunerRoot) TunerAdapter {
	return TunerAdapter{tuner.GetServicePort()}
}

// Tuner -> Ui
func (adapter *TunerAdapter) UpdateStationList(stationList domain.StationList) {
	httpPut(MakeUIEndpoint(UIStations), stationList)
}

// Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription domain.Subscription) {
	httpPut(MakeUIEndpoint(UISubscription), subscription)
}

// Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID domain.StationId) {
	httpPut(MakeHwEndpoint(HwCurrentStation), stationID)
}

func (adapter *TunerAdapter) handleSubscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var subscription domain.Subscription
		if err := decodeBody(r, &subscription); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.SubscriptionUpdated(subscription)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

func (adapter *TunerAdapter) handleStationList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var stations domain.StationList
		if err := decodeBody(r, &stations); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.StationListUpdated(stations)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

func (adapter *TunerAdapter) handleCurrentStation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var currentStationID uint32
		if err := decodeBody(r, &currentStationID); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.TuneToStation(currentStationID)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for Tuner
func (adapter *TunerAdapter) RunHTTPServer() {
	addr := TunerAddr
	mux := http.NewServeMux()
	mux.HandleFunc(TunerCurrentStations, adapter.handleCurrentStation)
	mux.HandleFunc(TunerStationList, adapter.handleStationList)
	mux.HandleFunc(TunerSubscription, adapter.handleSubscription)
	log.Println("Starting TunerAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping TunerAdapterr...")
}
