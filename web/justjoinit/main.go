package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// EmploymentType represents a salary/contract type (e.g. B2B, permanent) with a salary range.
type EmploymentType struct {
	From           *float64 `json:"from"`
	FromPerUnit    *float64 `json:"fromPerUnit"`
	To             *float64 `json:"to"`
	ToPerUnit      *float64 `json:"toPerUnit"`
	Currency       string   `json:"currency"`
	CurrencySource string   `json:"currencySource"`
	Type           string   `json:"type"`
	Unit           string   `json:"unit"`
	Gross          bool     `json:"gross"`
}

// Location represents a job offer's geographic location.
type Location struct {
	Slug      string  `json:"slug"`
	City      string  `json:"city"`
	Street    string  `json:"street"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Skill represents a skill with a name and proficiency level.
type Skill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Category represents a job category with an optional parent category.
type Category struct {
	Key       string  `json:"key"`
	ParentKey *string `json:"parentKey"`
}

// Offer represents a job offer from the JustJoin API.
type Offer struct {
	GUID                   string           `json:"guid"`
	Slug                   string           `json:"slug"`
	Title                  string           `json:"title"`
	WorkplaceType          string           `json:"workplaceType"`
	WorkingTime            string           `json:"workingTime"`
	ExperienceLevel        string           `json:"experienceLevel"`
	Category               Category         `json:"category"`
	City                   string           `json:"city"`
	Street                 string           `json:"street"`
	Latitude               float64          `json:"latitude"`
	Longitude              float64          `json:"longitude"`
	IsRemoteInterview      bool             `json:"isRemoteInterview"`
	CompanyName            string           `json:"companyName"`
	CompanyLogoURL         string           `json:"companyLogoThumbUrl"`
	PublishedAt            time.Time        `json:"publishedAt"`
	IsOpenToHireUkrainians bool             `json:"isOpenToHireUkrainians"`
	Locations              []Location       `json:"locations"`
	EmploymentTypes        []EmploymentType `json:"employmentTypes"`
	RequiredSkills         []Skill          `json:"requiredSkills"`
	NiceToHaveSkills       []Skill          `json:"niceToHaveSkills"`
	IsPromoted             bool             `json:"isPromoted"`
	IsSuperOffer           bool             `json:"isSuperOffer"`
	ApplyMethod            string           `json:"applyMethod"`
	ApplyUrl               *string          `json:"applyUrl"`
	LastPublishedAt        time.Time        `json:"lastPublishedAt"`
	ExpiredAt              time.Time        `json:"expiredAt"`
	OfferURL               string           `json:"-"` // actual offer URL
}

// APICursor represents pagination cursor with item count.
type APICursor struct {
	Cursor     *int `json:"cursor"`
	ItemsCount int   `json:"itemsCount"`
}

// APIMeta represents metadata for API pagination (from, total, prev/next cursor).
type APIMeta struct {
	From       int        `json:"from"`
	TotalItems int        `json:"totalItems"`
	Prev       *APICursor `json:"prev"`
	Next       *APICursor `json:"next"`
}

// APIResponse represents the response from the job offers API containing data and metadata.
type APIResponse struct {
	Data []Offer `json:"data"`
	Meta APIMeta `json:"meta"`
}

const (
	baseAPIURL = "https://justjoin.it/api/candidate-api/offers"
	pageSize   = 100
)

// formatSalary formats employment types into a human-readable salary string.
// Returns "N/A" if no salary info is available.
func formatSalary(types []EmploymentType) string {
	if len(types) == 0 {
		return "N/A"
	}
	for _, t := range types {
		if t.From != nil && t.To != nil {
			return fmt.Sprintf("%.0f-%.0f %s", *t.From, *t.To, t.Currency)
		}
		if t.From != nil {
			return fmt.Sprintf("from %.0f %s", *t.From, t.Currency)
		}
		if t.To != nil {
			return fmt.Sprintf("up to %.0f %s", *t.To, t.Currency)
		}
	}
	return "N/A"
}

// skillNames extracts skill names from a skill slice.
func skillNames(skills []Skill) []string {
	names := make([]string, len(skills))
	for i, s := range skills {
		names[i] = s.Name
	}
	return names
}

// fetchPage fetches a page of job offers from the API with given offset.
func fetchPage(from int) (*APIResponse, error) {
	params := url.Values{
		"categories": {"go"},
		"sortBy":     {"publishedAt"},
		"orderBy":    {"descending"},
		"from":       {fmt.Sprintf("%d", from)},
		"itemsCount": {fmt.Sprintf("%d", pageSize)},
	}

	req, err := http.NewRequest(http.MethodGet, baseAPIURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body[:min(len(body), 200)]))
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return &apiResp, nil
}

// FetchAllOffers fetches all job offers by paginating through the API.
func FetchAllOffers() ([]Offer, error) {
	var allOffers []Offer
	from := 0

	for {
		page, err := fetchPage(from)
		if err != nil {
			return nil, fmt.Errorf("fetching from offset %d: %w", from, err)
		}

		if len(page.Data) == 0 {
			break
		}

		allOffers = append(allOffers, page.Data...)

		if page.Meta.Next == nil || page.Meta.Next.Cursor == nil {
			break
		}
		if from+len(page.Data) >= page.Meta.TotalItems {
			break
		}
		from = *page.Meta.Next.Cursor
	}

	for i := range allOffers {
		allOffers[i].OfferURL = "https://justjoin.it/job-offer/" + allOffers[i].Slug
	}

	return allOffers, nil
}

// main is the entry point - fetches and prints all job offers.
func main() {
	offers, err := FetchAllOffers()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Fetched %d offers\n\n", len(offers))
	for i, o := range offers {
		fmt.Printf("%d. %s\n", i+1, o.Title)
		fmt.Printf("   - Company: %s\n", o.CompanyName)
		fmt.Printf("   - Salary: %s\n", formatSalary(o.EmploymentTypes))
		fmt.Printf("   - Mode: %s\n", o.WorkplaceType)
		if o.City != "" {
			fmt.Printf("   - Location: %s\n", o.City)
		} else {
			fmt.Printf("   - Location: N/A\n")
		}
		skills := skillNames(o.RequiredSkills)
		fmt.Printf("   - Technologies: %s\n", strings.Join(skills, ", "))
		fmt.Printf("   - Link: %s\n", o.OfferURL)
		fmt.Println()
	}
}
