package main

import (
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"log"
	"net/http"
	"rest"
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
	rest.HttpPut(rest.MakeUIEndpoint(rest.UIStations), stationList)
}

// Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription domain.Subscription) {
	rest.HttpPut(rest.MakeUIEndpoint(rest.UISubscription), subscription)
}

// Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID domain.StationId) {
	rest.HttpPut(rest.MakeHwEndpoint(rest.HwCurrentStation), stationID)
}

func (adapter *TunerAdapter) handleSubscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var subscription domain.Subscription
		if err := rest.DecodeBody(r, &subscription); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.SubscriptionUpdated(subscription)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

func (adapter *TunerAdapter) handleStationList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var stations domain.StationList
		if err := rest.DecodeBody(r, &stations); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.StationListUpdated(stations)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

func (adapter *TunerAdapter) handleCurrentStation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var currentStationID uint32
		if err := rest.DecodeBody(r, &currentStationID); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.TuneToStation(currentStationID)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for Tuner
func (adapter *TunerAdapter) RunHTTPServer() {
	addr := rest.TunerAddr
	mux := http.NewServeMux()
	mux.HandleFunc(rest.TunerCurrentStations, adapter.handleCurrentStation)
	mux.HandleFunc(rest.TunerStationList, adapter.handleStationList)
	mux.HandleFunc(rest.TunerSubscription, adapter.handleSubscription)
	log.Println("Starting TunerAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping TunerAdapterr...")
}
