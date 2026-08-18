// Command demo prints the latest Go job offers from JustJoin.it to stdout.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"findjobsmcp/api"
	"findjobsmcp/justjoinit"
)

func main() {
	jsonOut := flag.Bool("json", false, "print offers as formatted JSON")
	flag.Parse()

	offers, err := justjoinit.FetchAllOffers([]string{"go"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if *jsonOut {
		printOffersJSON(offers)
		return
	}

	fmt.Printf("Fetched %d offers\n\n", len(offers.Jobs))
	printOffers(offers.Jobs)
}

// printOffersJSON prints the given offers as indented JSON to standard output.
func printOffersJSON(offers api.EldoradoOffers) {
	out, err := json.MarshalIndent(offers, "", "  ")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(string(out))
}

// printOffers prints the given job offers to standard output.
func printOffers(offers []api.EldoradoOffer) {
	for i, o := range offers {
		fmt.Printf("%d. %s\n", i+1, o.Title)
		fmt.Printf("   - Company: %s\n", o.Company)
		fmt.Printf("   - Salary: %s\n", salary(o))
		fmt.Printf("   - Mode: %s\n", strings.Join(o.WorkModes, ", "))
		if len(o.Cities) > 0 {
			fmt.Printf("   - Location: %s\n", strings.Join(o.Cities, ", "))
		} else {
			fmt.Printf("   - Location: N/A\n")
		}
		fmt.Printf("   - Technologies: %s\n", strings.Join(o.Keywords, ", "))
		fmt.Printf("   - Link: %s\n", o.URL)
		fmt.Println()
	}
}

// salary formats an Eldorado salary range into a human-readable string.
func salary(o api.EldoradoOffer) string {
	if o.SalaryFrom != nil && o.SalaryTo != nil {
		return fmt.Sprintf("%d-%d PLN", *o.SalaryFrom, *o.SalaryTo)
	}
	if o.SalaryFrom != nil {
		return fmt.Sprintf("from %d PLN", *o.SalaryFrom)
	}
	if o.SalaryTo != nil {
		return fmt.Sprintf("up to %d PLN", *o.SalaryTo)
	}
	return "N/A"
}
