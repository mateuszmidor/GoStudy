package main

import (
	"net/http"
)

func main() {
	runWEB()
}

func runWEB() {
	http.Handle("/", NewTemplateHandler("map.html"))
	http.Handle("/api/find/json", NewFindRequestHandler())
	http.ListenAndServe(":9000", nil)
}
