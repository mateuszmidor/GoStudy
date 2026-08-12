package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

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

type Location struct {
	Slug      string  `json:"slug"`
	City      string  `json:"city"`
	Street    string  `json:"street"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Offer struct {
	ApplyMethod          string           `json:"applyMethod"`
	ApplyUrl             *string          `json:"applyUrl"`
	Body                 string           `json:"body"`
	CategoryID           int              `json:"categoryId"`
	City                 string           `json:"city"`
	CompanyLogoURL       string           `json:"companyLogoThumbUrl"`
	CompanyName          string           `json:"companyName"`
	CustomConsent        json.RawMessage  `json:"customConsent"`
	DisplayOffer         bool             `json:"displayOffer"`
	EmploymentTypes      []EmploymentType `json:"employmentTypes"`
	ExperienceLevel      string           `json:"experienceLevel"`
	ExpiredAt            time.Time        `json:"expiredAt"`
	FutureConsent        json.RawMessage  `json:"futureConsent"`
	GUID                 string           `json:"guid"`
	InformationClause    string           `json:"informationClause"`
	IsPromoted           bool             `json:"isPromoted"`
	IsSuperOffer         bool             `json:"isSuperOffer"`
	LastPublishedAt      time.Time        `json:"lastPublishedAt"`
	Latitude             float64          `json:"latitude"`
	Longitude            float64          `json:"longitude"`
	MatchPercent         string           `json:"matchPercent"`
	Multilocation        []Location       `json:"multilocation"`
	NiceToHaveSkills     []string         `json:"niceToHaveSkills"`
	OpenToHireUkrainians bool             `json:"openToHireUkrainians"`
	PublishedAt          time.Time        `json:"publishedAt"`
	RemoteInterview      bool             `json:"remoteInterview"`
	RequiredSkills       []string         `json:"requiredSkills"`
	Slug                 string           `json:"slug"`
	Details              string           `json:"-"` // actual offer details link
	Street               string           `json:"street"`
	Title                string           `json:"title"`
	WorkingTime          string           `json:"workingTime"`
	WorkplaceType        string           `json:"workplaceType"`
}

const offersURL = "https://justjoin.it/job-offers/all-locations/go?sortBy=published"

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

var scriptPushRE = regexp.MustCompile(`self\.__next_f\.push\(\[1,"((?:[^"\\]|\\.)*)"\]\)`)

func FetchOffers(url string) ([]Offer, error) {
	if url == "" {
		url = offersURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	jsonData, err := extractRSCJSON(string(body))
	if err != nil {
		return nil, fmt.Errorf("extracting JSON: %w", err)
	}

	offers, err := parseOffers(jsonData)
	if err != nil {
		return nil, fmt.Errorf("parsing offers: %w", err)
	}

	return offers, nil
}

func extractRSCJSON(html string) (string, error) {
	matches := scriptPushRE.FindAllStringSubmatch(html, -1)
	if matches == nil {
		return "", fmt.Errorf("no RSC script data found")
	}

	for _, m := range matches {
		inner := m[1]
		if strings.Contains(inner, "OFFERS") && strings.Contains(inner, "requiredSkills") {
			unescaped := strings.ReplaceAll(inner, `\"`, `"`)
			unescaped = strings.ReplaceAll(unescaped, `\\`, `\`)
			unescaped = strings.ReplaceAll(unescaped, `\n`, "\n")
			return unescaped, nil
		}
	}

	return "", fmt.Errorf("no OFFERS payload found in RSC data")
}

func parseOffers(jsonStr string) ([]Offer, error) {
	// RSC payload format: XX:["$","$L39",null,{...}]
	// Find the first '{' which starts the inner JSON object
	idx := strings.Index(jsonStr, "null,{")
	if idx < 0 {
		return nil, fmt.Errorf("cannot find payload object in RSC data")
	}
	jsonStr = jsonStr[idx+5:] // skip "null,"

	type page struct {
		Data json.RawMessage `json:"data"`
		Meta json.RawMessage `json:"meta"`
	}

	var root struct {
		State struct {
			Queries []struct {
				QueryKey []interface{} `json:"queryKey"`
				State    struct {
					Data json.RawMessage `json:"data"`
				} `json:"state"`
			} `json:"queries"`
		} `json:"state"`
	}

	dec := json.NewDecoder(strings.NewReader(jsonStr))
	if err := dec.Decode(&root); err != nil {
		return nil, fmt.Errorf("unmarshaling RSC JSON: %w", err)
	}

	for _, q := range root.State.Queries {
		keyStr := fmt.Sprintf("%v", q.QueryKey)
		if !strings.Contains(keyStr, "OFFERS") || strings.Contains(keyStr, "COUNT") || strings.Contains(keyStr, "BANNER") {
			continue
		}

		var data struct {
			Pages []page `json:"pages"`
		}
		if err := json.Unmarshal(q.State.Data, &data); err != nil {
			return nil, fmt.Errorf("unmarshaling pages data: %w", err)
		}
		if len(data.Pages) == 0 {
			return nil, fmt.Errorf("no pages in OFFERS data")
		}

		var offers []Offer
		if err := json.Unmarshal(data.Pages[0].Data, &offers); err != nil {
			return nil, fmt.Errorf("unmarshaling offers array: %w", err)
		}
		for i := range offers {
			offers[i].Details = "https://justjoin.it/job-offer/" + offers[i].Slug
		}
		return offers, nil
	}

	return nil, fmt.Errorf("OFFERS query not found")
}

func main() {
	offers, err := FetchOffers("")
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
		fmt.Printf("   - Technologies: %s\n", strings.Join(o.RequiredSkills, ", "))
		fmt.Printf("   - Link: %s\n", o.Details)
		fmt.Println()
	}
}
