package main

import (
	"fmt"
	"log"
	"strings"
)

const baseURL = "https://edu.devstyle.pl"

func main() {
	client, err := newAuthorizedClient(baseURL+"/login", email, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()

	// Collect all categories from pages 1–3.
	categoriesBase := baseURL + "/products/dna-droga-nowoczesnego-architekta/categories"
	var categories []SyllabusItem
	for page := 1; page <= 3; page++ {
		pageURL := fmt.Sprintf("%s?page=%d", categoriesBase, page)
		items, err := fetchSyllabusItems(client, pageURL)
		if err != nil {
			log.Fatalln(err)
		}
		categories = append(categories, items...)
	}

	// For each category, fetch its posts and print the tree.
	for _, cat := range categories {
		fmt.Println(cat.Title)

		// Links may be relative (/products/...) or absolute (https://...).
		catURL := cat.Link
		if !strings.HasPrefix(catURL, "http") {
			catURL = baseURL + catURL
		}
		posts, err := fetchSyllabusItems(client, catURL)
		if err != nil {
			log.Fatalln(err)
		}
		for _, post := range posts {
			postURL := post.Link
			if !strings.HasPrefix(postURL, "http") {
				postURL = baseURL + postURL
			}
			fmt.Printf("  - %s\n    %s\n", post.Title, postURL)
		}
		fmt.Println()
	}
}
