// Project: Marshall go data into JSON string
// Usage: go run .
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie describes movie meta data
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Nieugięty Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	const prefix = ""   // first char of a line
	const indent = "  " // for block indentation

	// struct -> JSON
	jsonBytes, err := json.MarshalIndent(movies, prefix, indent)
	if err != nil {
		log.Fatalf("Marshalling to JSON failed: %s", err)
	}
	fmt.Printf("%s\n", jsonBytes)

	// JSON -> full struct
	var wholeMovies []Movie
	if err := json.Unmarshal(jsonBytes, &wholeMovies); err != nil {
		log.Fatalf("Unmarshalling from JSON failed :%s", err)
	}
	fmt.Printf("%v\n", wholeMovies)

	// JSON -> Title only struct
	var onlyTitles []struct{ Title string } // only Title field from json will be read
	if err := json.Unmarshal(jsonBytes, &onlyTitles); err != nil {
		log.Fatalf("Unmarshalling from JSON failed :%s", err)
	}
	fmt.Printf("%v\n", onlyTitles)
}
