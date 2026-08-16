package main

import (
	"fmt"

	"findjobsmcp/justjoinit"
)

// main is the entry point - fetches Go job offers and prints them.
func main() {
	offers, err := justjoinit.FetchAllOffers([]string{"go"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Fetched %d offers\n\n", len(offers.Jobs))
	printOffers(offers.Jobs)
}