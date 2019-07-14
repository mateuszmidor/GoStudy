package main

import (
	"hexagons/ui"
	"hexagons/ui/infrastructure"
	"log"
	"net/http"
	"rest"
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
	rest.HttpPut(rest.MakeTunerEndpoint(rest.TunerCurrentStations), stationID)
}

func (adapter *UIAdapter) handleStations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var stations []string
		if err := rest.DecodeBody(r, &stations); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read stations from request", err)
			return
		}

		adapter.uiServicePort.UpdateStationList(stations)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

func (adapter *UIAdapter) handleSubscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var subscription bool
		if err := rest.DecodeBody(r, &subscription); err != nil {
			rest.RespondErr(w, http.StatusBadRequest, "Couldnt read subscription status from request", err)
			return
		}

		adapter.uiServicePort.UpdateSubscription(subscription)
		rest.Respond(w, http.StatusOK, nil)
		return
	}
	rest.RespondHTTPErr(w, http.StatusNotFound)
}

// RunHttpServer starts a server that handles commands for UI
func (adapter *UIAdapter) RunHTTPServer() {
	addr := rest.UIAddr
	mux := http.NewServeMux()
	mux.HandleFunc(rest.UIStations, adapter.handleStations)
	mux.HandleFunc(rest.UISubscription, adapter.handleSubscription)
	log.Println("Starting UIAdapter at", addr)
	http.ListenAndServe(addr, mux)
	log.Println("Stopping UIAdapterr...")
}
