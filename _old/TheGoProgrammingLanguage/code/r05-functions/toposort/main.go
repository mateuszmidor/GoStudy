// Project: find correct order to take courses so you have all prerequisites necessary to take them all
// Usage: go run .
package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequsites
var preqs = map[string][]string{
	"algorytmy":                      {"struktury danych"},
	"rachunek rozniczkowy i calkowy": {"algebra liniowa"},
	"kompilatory":                    {"struktury danych", "jezyki formalne", "organizacja procesora"},
	"struktury danych":               {"matematyka dyskretna"},
	"bazy danych":                    {"struktury danych"},
	"matematyka dyskretna":           {"wstep do programowania"},
	"jezyki formalne":                {"matematyka dyskretna"},
	"sieci":                          {"systemy operacyjne"},
	"systemy operacyjne":             {"strutury danych", "organizacja procesora"},
	"jezyki programowania":           {"strutury danych", "organizacja procesora"},
}

func main() {
	for i, course := range topoSort(preqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
