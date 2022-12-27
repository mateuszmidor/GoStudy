package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

var degrees float64 = 0.0

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	value := math.Sin(degrees * math.Pi / 180.0)
	metric := fmt.Sprintf("wave %f", value)

	fmt.Printf("Metric requested. Returning:\n%s\n\n", metric)
	fmt.Fprintf(w, metric)

	degrees += 15.0
}

func main() {
	fmt.Println("Running metrics provider at localhost:8080/metrics")

	http.HandleFunc("/metrics", handleMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
