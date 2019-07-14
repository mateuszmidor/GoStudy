package main

import (
	"hexagons/ui"
	"hexagons/ui/infrastructure"
	"log"
	"net/http"
)

// UIAdapter implements tuner output ports towards ui, and ui output ports towards tuner
type UIAdapter struct {
	uiServicePort infrastructure.ServicePort
}

// NewUIAdapter creates a HTTP adapter for UI
func NewUIAdapter(ui *ui.UiRoot) UIAdapter {
	return UIAdapter{ui.GetServicePort()}
}

// TuneToStation forwards command UI -> Tuner
func (adapter *UIAdapter) TuneToStation(stationID uint32) {
	httpPut(MakeTunerEndpoint(TunerCurrentStations), stationID)
}

func (adapter *UIAdapter) handleStations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var stations []string
		if err := decodeBody(r, &stations); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read stations from request", err)
			return
		}

		adapter.uiServicePort.UpdateStationList(stations)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

func (adapter *UIAdapter) handleSubscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var subscription bool
		if err := decodeBody(r, &subscription); err != nil {
			respondErr(w, http.StatusBadRequest, "Couldnt read subscription status from request", err)
			return
		}

		adapter.uiServicePort.UpdateSubscription(subscription)
		respond(w, http.StatusOK, nil)
		return
	}
	respondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for UI
func (adapter *UIAdapter) RunHTTPServer() {
	addr := UIAddr
	mux := http.NewServeMux()
	mux.HandleFunc(UIStations, adapter.handleStations)
	mux.HandleFunc(UISubscription, adapter.handleSubscription)
	log.Println("Starting UIAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping UIAdapterr...")
}
