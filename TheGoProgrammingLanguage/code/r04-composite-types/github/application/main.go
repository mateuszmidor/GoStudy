// Project: Query github for issues and print the out
// Usage: go run .
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mateuszmidor/GoStudy/TheGoProgrammingLanguage/code/r04/github/githubapi"
)

func main() {
	oneMonthAgo := time.Now().AddDate(0, -1, 0).Format("2006-01-02")
	terms := []string{"repo:golang/go", "is:open", "created:>=" + oneMonthAgo, "state:open", "json", "decoder"}

	result, err := githubapi.SearchIssues(terms)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d topics created in last 30 days:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%05d %10v %9.9s %.100s\n", item.Number, item.CreatedAt.Format("2006-01-02"), item.User.Login, item.Title)
	}
}
