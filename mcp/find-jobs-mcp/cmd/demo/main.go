// Command demo prints job offers from JustJoin.it to stdout.
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
	category := flag.String("category", "", "job category: "+api.OfferCategoriesStr)
	keywords := flag.String("keywords", "", "full-text search, e.g. golang, react developer")
	city := flag.String("city", "", "filter by city, e.g. Warszawa, Kraków")
	experienceLevels := flag.String("experience-levels", "", "comma-separated: intern, junior, mid, senior, manager, c_level")
	employmentTypes := flag.String("employment-types", "", "comma-separated: b2b, permanent, uoz, internship")
	remoteWorkOptions := flag.String("work-modes", "", "comma-separated: remote, hybrid, office")
	withSalary := flag.Bool("with-salary", false, "only offers with disclosed salary")
	minSalary := flag.Int("min-salary", 0, "minimum salary threshold in PLN")
	sortBy := flag.String("sort-by", "publishedAt", "sort field: publishedAt, salary")
	orderBy := flag.String("order-by", "descending", "sort direction: ascending, descending")
	jsonOut := flag.Bool("json", false, "print offers as formatted JSON")
	flag.Parse()

	params := justjoinit.SearchParams{}
	params.Categories = []string{*category}
	params.Keywords = *keywords
	params.City = *city
	params.ExperienceLevels = splitTrim(*experienceLevels)
	params.EmploymentTypes = splitTrim(*employmentTypes)
	params.RemoteWorkOptions = splitTrim(*remoteWorkOptions)
	params.WithSalary = withSalary
	params.MinSalary = minSalary
	params.SortBy = *sortBy
	params.OrderBy = *orderBy

	fmt.Printf("Listing offers for %+v\n", params)
	offers, err := justjoinit.FetchAllOffers(params)
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

// splitTrim splits a comma-separated string and trims whitespace from each part.
func splitTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
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
