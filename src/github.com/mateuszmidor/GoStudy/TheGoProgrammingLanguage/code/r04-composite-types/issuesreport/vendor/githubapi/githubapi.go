package githubapi

import (
	"time"

	githubapi "github.com/mateuszmidor/GoStudy/TheGoProgrammingLanguage/code/r04-composite-types/github/githubapi"
)

// SsearchIssues querries GitHub issues API
func searchIssues(terms []string) (*githubapi.IssuesSearchResult, error) {
	return githubapi.SearchIssues(terms)
}

// SearchIssuesSince3Months looks up github issues up to 3 months old
func SearchIssuesSince3Months() (*githubapi.IssuesSearchResult, error) {
	threeMonthsAgo := time.Now().AddDate(0, -3, 0).Format("2006-01-02")
	terms := []string{"repo:golang/go", "is:open", "created:>=" + threeMonthsAgo, "state:open", "json", "decoder"}
	return searchIssues(terms)
}
