// Project: 3d plotter
// Usage: go run . & firefox localhost:8000/formula
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/formula", handleFormulaPage)
	http.HandleFunc("/plot", handlePlotPage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleFormulaPage(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "formula.html")
}

func handlePlotPage(w http.ResponseWriter, req *http.Request) {
	formula := req.URL.Query().Get("formula")
	fmt.Printf("formula(x, y) = %s\n", formula)

	w.Header().Set("Content-Type", "image/svg+xml")
	if err := plotSVG3D(w, formula); err != nil {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, err)
	}
}
