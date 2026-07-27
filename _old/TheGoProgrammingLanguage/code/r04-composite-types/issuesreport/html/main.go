// Project: print GitHub issues using html/template package
// html/template automatically escapes the template contents so it is not "executed" by browser by displayed as-is
// This is to prevent code injection attacks
// Usage: go run .
package main

import (
	"githubapi"
	"html/template"
	"log"
	"os"
	"time"
)

var templ = `
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<h1>Found issues from last three months - {{.TotalCount}}:</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
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
