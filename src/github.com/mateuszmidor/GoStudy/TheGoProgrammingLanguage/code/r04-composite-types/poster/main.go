// Project: Lookup movie and return it's poster URL
// Usage: go run . "forrest gump"
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const apiKey = "d23cadab"
const requestPattern = "http://www.omdbapi.com/?t=%s&apikey=%s"

type omdbMovie struct {
	Title    string
	Year     string
	Runtime  string
	Actors   string
	Plot     string
	Poster   string
	Response string
}

func main() {
	movieRequest := getHTTPMovieRequest()
	movie := getOMDBMovie(movieRequest)
	printMovie(movie)
}

func getHTTPMovieRequest() string {
	if len(os.Args) != 2 {
		fmt.Printf("Please specify movie title, eg. %q", "Matrix")
		os.Exit(1)
	}

	movieTitle := os.Args[1]
	escaptedTile := url.QueryEscape(movieTitle)
	movieRequest := fmt.Sprintf(requestPattern, escaptedTile, apiKey)
	return movieRequest
}

func getOMDBMovie(url string) omdbMovie {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var result omdbMovie
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return result
}

func printMovie(m omdbMovie) {
	if m.Response == "True" { // movie found
		fmt.Printf("%-10s: %s\n", "Title", m.Title)
		fmt.Printf("%-10s: %s\n", "Year", m.Year)
		fmt.Printf("%-10s: %s\n", "Runtime", m.Runtime)
		fmt.Printf("%-10s: %s\n", "Actors", m.Actors)
		fmt.Printf("%-10s: %s\n", "Plot", m.Plot)
		fmt.Printf("%-10s: %s\n", "Poster", m.Poster)
	} else {
		fmt.Println("Movie not found!")
	}
}
