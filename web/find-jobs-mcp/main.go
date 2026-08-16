package main

import (
	"fmt"
	"strings"

	"findjobsmcp/api"
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
	for i, o := range offers.Jobs {
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
