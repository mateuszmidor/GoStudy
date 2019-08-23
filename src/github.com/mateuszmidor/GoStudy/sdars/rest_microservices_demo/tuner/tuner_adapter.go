package main

import (
	"hexagons/tuner"
	"hexagons/tuner/domain"
	"hexagons/tuner/infrastructure"
	"log"
	"net/http"
	"rest"
	"retry"
)

type TunerAdapter struct {
	tunerServicePort infrastructure.TunerServicePort
}

// NewTunerAdapter creates a HTTP adapter for Tuner hexagon
func NewTunerAdapter(tuner *tuner.TunerRoot) TunerAdapter {
	return TunerAdapter{tuner.GetServicePort()}
}

// Tuner -> Ui
func (adapter *TunerAdapter) UpdateStationList(stationList domain.StationList) {
	endpoint := rest.MakeUIEndpoint(rest.UIStations)
	retry.UntilSuccessOr5Failures("updating station list", rest.HttpPut, endpoint, stationList)
}

// Tuner -> Ui
func (adapter *TunerAdapter) UpdateSubscription(subscription domain.Subscription) {
	endpoint := rest.MakeUIEndpoint(rest.UISubscription)
	retry.UntilSuccessOr5Failures("updating subscription", rest.HttpPut, endpoint, subscription)
}

// Tuner -> Hw
func (adapter *TunerAdapter) TuneToStation(stationID domain.StationID) {
	endpoint := rest.MakeHwEndpoint(rest.HwCurrentStation)
	retry.UntilSuccessOr5Failures("tuning to station", rest.HttpPut, endpoint, stationID)
}

func (adapter *TunerAdapter) handleSubscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var subscription domain.Subscription
		if err := rest.DecodeBody(r, &subscription); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read current station id from request", err)
			return
		}
		adapter.tunerServicePort.UpdateSubscription(subscription)
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
		adapter.tunerServicePort.UpdateStationList(stations)
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
