// Project: print GitHub issues using text/template package
// Usage: go run .
package main

import (
	"githubapi"
	"log"
	"os"
	"text/template"
	"time"
)

var templ = `
Found issues from last 3 months - {{.TotalCount}}:
{{range .Items}}-------------------------------------------
Number:         {{.Number}}
User:           {{.User.Login}}
Title:          {{.Title | printf "%.64s"}}
Created:        {{.CreatedAt | daysAgo}} days ago 
{{end}}
`

func main() {
	funcMap := template.FuncMap{"daysAgo": daysAgo} // will handle "| daysAgo" from template
	report := template.Must(template.New("issuelist").Funcs(funcMap).Parse(templ))

	issues, err := githubapi.SearchIssuesSince3Months()
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, issues); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
