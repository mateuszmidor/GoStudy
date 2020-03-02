// Package githubapi allows to look up issues on GitHub.
// It is not vendored so can be imported and used by other examples from the book
// Example github api request: https://api.github.com/search/issues?q=repo:golang/go+is:open+json
package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL is github issue search endpoint
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is complete search result boundle fetched from github
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is single github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User is issue author
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues querries GitHub issues API
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}
