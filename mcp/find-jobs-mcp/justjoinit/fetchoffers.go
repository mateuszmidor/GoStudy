package justjoinit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"findjobsmcp/api"
)

const (
	baseAPIURL = "https://justjoin.it/api/candidate-api/offers"
	pageSize   = 100
)

// SearchParams holds optional filters for the JustJoin.it candidate API.
// All fields are optional; zero values are omitted from the query.
// Multiple filters are combined with AND logic.
type SearchParams struct {
	// Categories filters by job category slug.
	// Allowed values: "ai", "go", "html", "java", "python", "javascript", "php",
	// "ruby", "net", "scala", "c", "mobile", "testing", "devops", "admin",
	// "ux", "pm", "game", "analytics", "security", "data", "support",
	// "erp", "architecture".
	Categories []string

	// Keywords is a full-text search phrase matched against job titles
	// and required/nice-to-have skills, e.g. "golang", "react developer".
	Keywords string

	// City filters by city name (diacritics optional, e.g. "Krakow" or "Kraków").
	City string

	// ExperienceLevels filters by seniority level.
	// Allowed values: "intern", "junior", "mid", "senior", "manager", "c_level".
	ExperienceLevels []string

	// EmploymentTypes filters by contract type.
	// Allowed values: "b2b", "permanent", "uoz" (mandate contract),
	// "internship" (specific-task contract).
	EmploymentTypes []string

	// RemoteWorkOptions filters by workplace arrangement.
	// Allowed values: "remote", "hybrid", "office".
	RemoteWorkOptions []string

	// WithSalary when non-nil and true, returns only offers that disclose salary.
	WithSalary *bool

	// MinSalary when non-nil, sets a minimum salary threshold in PLN.
	MinSalary *int

	// SortBy controls the sort field. Default: "publishedAt".
	// Allowed values: "publishedAt", "salary".
	SortBy string

	// OrderBy controls the sort direction. Default: "descending".
	// Allowed values: "ascending", "descending".
	OrderBy string
}

// employmentType represents a salary/contract type (e.g. B2B, permanent) with a salary range.
type employmentType struct {
	From           *float64 `json:"from"`           // monthly amount in Currency, e.g. 26765.0
	FromPerUnit    *float64 `json:"fromPerUnit"`    // rate per Unit, e.g. 159.3146 (PLN/h when Unit is Hour)
	To             *float64 `json:"to"`             // monthly amount in Currency, e.g. 36169.0
	ToPerUnit      *float64 `json:"toPerUnit"`      // rate per Unit, e.g. 220.0
	Currency       string   `json:"currency"`       // e.g. "PLN" (also "USD", "EUR", "CHF", "GBP")
	CurrencySource string   `json:"currencySource"` // e.g. "original" (converted offers use "conversion")
	Type           string   `json:"type"`           // e.g. "b2b" (also "permanent", "any")
	Unit           string   `json:"unit"`           // e.g. "Month" or "Hour"
	Gross          bool     `json:"gross"`          // e.g. false
}

// location represents a job offer's geographic location.
type location struct {
	Slug      string  `json:"slug"`      // e.g. "asana-senior-software-engineer-platform-reliability-warszawa-go-8a93939c"
	City      string  `json:"city"`      // e.g. "Warszawa"
	Street    string  `json:"street"`    // e.g. "Marcina Kasprzaka 6"
	Latitude  float64 `json:"latitude"`  // e.g. 52.2296513
	Longitude float64 `json:"longitude"` // e.g. 20.9747806
}

// skill represents a skill with a name and proficiency level.
type skill struct {
	Name  string `json:"name"`  // e.g. "AWS"
	Level int    `json:"level"` // 1-5, e.g. 4
}

// category represents a job category with an optional parent category.
type category struct {
	Key       string  `json:"key"`       // e.g. "go"
	ParentKey *string `json:"parentKey"` // e.g. nil (null for all top-level categories like "go")
}

// offer represents a job offer from the JustJoin API.
type offer struct {
	GUID                   string           `json:"guid"`                   // e.g. "9459e8f2-d23c-48ee-92a1-3543262100e6"
	Slug                   string           `json:"slug"`                   // e.g. "asana-senior-software-engineer-platform-reliability-warszawa-go-8a93939c"
	Title                  string           `json:"title"`                  // e.g. "Senior Software Engineer, Platform Reliability"
	WorkplaceType          string           `json:"workplaceType"`          // e.g. "hybrid" (also "remote", "office", "partly_remote")
	WorkingTime            string           `json:"workingTime"`            // e.g. "full_time" (also "part_time")
	ExperienceLevel        string           `json:"experienceLevel"`        // e.g. "senior" (also "junior", "mid", "c_level")
	Category               category         `json:"category"`               // e.g. {key "go", parentKey nil}
	City                   string           `json:"city"`                   // e.g. "Warszawa", empty for remote-only offers
	Street                 string           `json:"street"`                 // e.g. "Marcina Kasprzaka 6"
	Latitude               float64          `json:"latitude"`               // e.g. 52.2296513
	Longitude              float64          `json:"longitude"`              // e.g. 20.9747806
	IsRemoteInterview      bool             `json:"isRemoteInterview"`      // e.g. false
	CompanyName            string           `json:"companyName"`            // e.g. "Asana"
	CompanyLogoURL         string           `json:"companyLogoThumbUrl"`    // e.g. "https://imgproxy.justjoinit.tech/..."
	PublishedAt            time.Time        `json:"publishedAt"`            // e.g. "2026-08-16T07:00:05.6295683Z"
	IsOpenToHireUkrainians bool             `json:"isOpenToHireUkrainians"` // e.g. false
	Locations              []location       `json:"locations"`              // e.g. [{City "Warszawa", Street "Centrum"}, ...]
	EmploymentTypes        []employmentType `json:"employmentTypes"`        // e.g. [{Type "b2b", Unit "Hour", From 180.0, To 220.0}]
	RequiredSkills         []skill          `json:"requiredSkills"`         // e.g. [{Name "AWS", Level 4}, {Name "Go", Level 4}]
	NiceToHaveSkills       []skill          `json:"niceToHaveSkills"`       // e.g. [] (often empty)
	IsPromoted             bool             `json:"isPromoted"`             // e.g. false
	IsSuperOffer           bool             `json:"isSuperOffer"`           // e.g. false
	ApplyMethod            string           `json:"applyMethod"`            // e.g. "external"
	ApplyURL               *string          `json:"applyUrl"`               // e.g. "https://grnh.se/dm05zac51us", nil when unavailable
	LastPublishedAt        time.Time        `json:"lastPublishedAt"`        // e.g. "2026-07-02T06:21:27.015058Z"
	ExpiredAt              time.Time        `json:"expiredAt"`              // e.g. "2026-08-31T21:59:59.999Z"
	OfferURL               string           `json:"-"`                      // filled by us, e.g. "https://justjoin.it/job-offer/<slug>"
}

// apiCursor represents pagination cursor with item count.
type apiCursor struct {
	Cursor     *int `json:"cursor"`     // e.g. 81 (nil when there are no more pages)
	ItemsCount int  `json:"itemsCount"` // e.g. 81
}

// apiMeta represents metadata for API pagination (from, total, prev/next cursor).
type apiMeta struct {
	From       int        `json:"from"`       // offset of the first item, e.g. 0
	TotalItems int        `json:"totalItems"` // total number of matching offers, e.g. 81
	Prev       *apiCursor `json:"prev"`       // e.g. {nil, 81} on the first page
	Next       *apiCursor `json:"next"`       // e.g. {81, 81} on the first page
}

// apiResponse represents the response from the job offers API containing data and metadata.
type apiResponse struct {
	Data []offer `json:"data"` // e.g. 100 offers
	Meta apiMeta `json:"meta"` // e.g. {From 0, TotalItems 81, Next {81, 81}}
}

// fetchPage fetches a page of job offers from the API with given offset and search params.
func fetchPage(params SearchParams, from int) (*apiResponse, error) {
	q := url.Values{}
	for _, c := range params.Categories {
		q.Add("categories", c)
	}
	for _, v := range params.ExperienceLevels {
		q.Add("experienceLevels", v)
	}
	for _, v := range params.EmploymentTypes {
		q.Add("employmentTypes", v)
	}
	for _, v := range params.RemoteWorkOptions {
		q.Add("remoteWorkOptions", v)
	}
	if params.Keywords != "" {
		q.Set("keywords", params.Keywords)
	}
	if params.City != "" {
		q.Set("city", params.City)
	}
	if params.WithSalary != nil && *params.WithSalary {
		q.Set("withSalary", "true")
	}
	if params.MinSalary != nil {
		q.Set("minSalary", fmt.Sprintf("%d", *params.MinSalary))
	}
	sortBy := params.SortBy
	if sortBy == "" {
		sortBy = "publishedAt"
	}
	orderBy := params.OrderBy
	if orderBy == "" {
		orderBy = "descending"
	}
	q.Set("sortBy", sortBy)
	q.Set("orderBy", orderBy)
	q.Set("from", fmt.Sprintf("%d", from))
	q.Set("itemsCount", fmt.Sprintf("%d", pageSize))

	req, err := http.NewRequest(http.MethodGet, baseAPIURL+"?"+q.Encode(), nil)
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

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	return &apiResp, nil
}

// FetchAllOffers fetches all job offers matching the given search parameters
// and translates them into the Eldorado representation. When params is empty
// (zero value), all offers are returned sorted by most recently published.
func FetchAllOffers(params SearchParams) (api.EldoradoOffers, error) {
	var allOffers []offer
	from := 0

	for {
		page, err := fetchPage(params, from)
		if err != nil {
			return api.EldoradoOffers{}, fmt.Errorf("fetching from offset %d: %w", from, err)
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

	return offersToEldorado(allOffers), nil
}
