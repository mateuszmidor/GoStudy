package api

import (
	"strings"
)

// OfferCategories are the searchable offer categories :)
var OfferCategories = []string{"ai", "go", "java", "python", "javascript", "php", "ruby", "net", "scala", "c", "mobile", "testing", "devops", "admin", "ux", "pm", "game", "analytics", "security", "data", "support", "erp", "architecture"}
var OfferCategoriesStr = strings.Join(OfferCategories, ", ")

// EldoradoOffers is the format returned from czyjesteldorado.pl MCP server
type EldoradoOffers struct {
	Jobs []EldoradoOffer `json:"jobs"` // e.g. 53 offers
}

// EldoradoOffer represents a single job offer from CzyJestEldorado.
type EldoradoOffer struct {
	Title           string   `json:"title"`            // e.g. "Senior Software Architect"
	Keywords        []string `json:"keywords"`         // e.g. ["Go", "Angular", "NATS", "Kafka"]
	Company         string   `json:"company"`          // e.g. "Michael Page"
	Cities          []string `json:"cities"`           // e.g. ["Warszawa", "Kraków"]; empty [] for fully remote offers
	WorkModes       []string `json:"work_modes"`       // allowed modes: "remote" | "hybrid" | "office", e.g. ["hybrid"]
	ContractTypes   []string `json:"contract_types"`   // allowed contracts: "b2b" | "employment_contract", may be empty
	EmploymentTypes []string `json:"employment_types"` // e.g. ["full_time"]
	Seniority       string   `json:"seniority"`        // one of: intern | junior | mid | senior | c_level
	SalaryFrom      *int     `json:"salary_from"`      // monthly minimum gross PLN, nil when undisclosed, e.g. 28000
	SalaryTo        *int     `json:"salary_to"`        // monthly maximum gross PLN, nil when undisclosed, e.g. 35000
	URL             string   `json:"url"`              // direct offer link, e.g. "https://czyjesteldorado.pl/praca/399188-..."
}
