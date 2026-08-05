package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// userAgent mimics a real browser to pass Cloudflare's bot detection.
// The full UA string (including AppleWebKit/Chrome tokens) is required;
// a shortened version is treated as a bot and blocked with 403.
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

const baseURL = "https://edu.devstyle.pl"

// SyllabusItem holds the title and relative link of a syllabus entry
// (used for both categories and posts within a category).
type SyllabusItem struct {
	Title string
	Link  string
}

// doGet sends a GET request with the browser User-Agent set.
// Every request must carry this header; without it Cloudflare returns 403.
func doGet(client *http.Client, targetURL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)
	return client.Do(req)
}

// extractCSRFToken parses HTML and returns the value of the first
// <input name="authenticity_token"> element (Rails/Kajabi CSRF field).
func extractCSRFToken(body io.Reader) string {
	doc, err := html.Parse(body)
	if err != nil {
		return ""
	}

	var traverse func(*html.Node) string
	traverse = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "input" {
			var name, value string
			for _, attr := range n.Attr {
				if attr.Key == "name" {
					name = attr.Val
				}
				if attr.Key == "value" {
					value = attr.Val
				}
			}
			if name == "authenticity_token" {
				return value
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if val := traverse(c); val != "" {
				return val
			}
		}
		return ""
	}

	return traverse(doc)
}

// firstHref returns the href of the first <a> element found under n.
func firstHref(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if href := firstHref(c); href != "" {
			return href
		}
	}
	return ""
}

// syllabusTitle returns the trimmed text of the first
// <p class="syllabus__title"> element found under n.
func syllabusTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "p" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "syllabus__title" {
				var buf strings.Builder
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.TextNode {
						buf.WriteString(strings.TrimSpace(c.Data))
					}
				}
				return buf.String()
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if t := syllabusTitle(c); t != "" {
			return t
		}
	}
	return ""
}

// extractSyllabusItems finds all <div class="syllabus__item"> blocks and
// returns each one's title (from <p class="syllabus__title">) and link
// (from the first <a href> inside the block).
func extractSyllabusItems(n *html.Node, items *[]SyllabusItem) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == "syllabus__item" {
				title := syllabusTitle(n)
				href := firstHref(n)
				if title != "" {
					*items = append(*items, SyllabusItem{Title: title, Link: href})
				}
				return // don't recurse into nested syllabus__item blocks
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractSyllabusItems(c, items)
	}
}

// fetchSyllabusItems fetches a URL and returns all syllabus items on that page.
func fetchSyllabusItems(client *http.Client, pageURL string) ([]SyllabusItem, error) {
	resp, err := doGet(client, pageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var items []SyllabusItem
	extractSyllabusItems(doc, &items)
	return items, nil
}

func main() {
	loginURL := "https://edu.devstyle.pl/login"

	// 1. Create a cookie jar so the client automatically handles session cookies.
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{Jar: jar}

	// 2. GET the login page to obtain session cookies and the CSRF token.
	resp, err := doGet(client, loginURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("GET login page status:", resp.Status)

	csrfToken := extractCSRFToken(resp.Body)
	fmt.Println("CSRF token found:", csrfToken != "")

	// 3. Build the login form payload.

	formData := url.Values{
		"utf8":                {"✓"},
		"member[email]":       {email},
		"member[password]":    {password},
		"member[remember_me]": {"1"},
		"commit":              {"Login »"},
	}
	if csrfToken != "" {
		formData.Set("authenticity_token", csrfToken)
	}

	// 4. POST credentials to log in.
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", loginURL)

	loginResp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer loginResp.Body.Close()
	fmt.Println("Login response status:", loginResp.Status)
	fmt.Println("Login final URL:", loginResp.Request.URL.String())
	fmt.Println()

	// 5. Collect all categories from pages 1–3.
	categoriesBase := baseURL + "/products/dna-droga-nowoczesnego-architekta/categories"
	var categories []SyllabusItem
	for page := 1; page <= 3; page++ {
		pageURL := fmt.Sprintf("%s?page=%d", categoriesBase, page)
		items, err := fetchSyllabusItems(client, pageURL)
		if err != nil {
			panic(err)
		}
		categories = append(categories, items...)
	}

	// 6. For each category, fetch its posts and print the tree.
	for _, cat := range categories {
		fmt.Println(cat.Title)

		// Links may be relative (/products/...) or absolute (https://...).
		catURL := cat.Link
		if !strings.HasPrefix(catURL, "http") {
			catURL = baseURL + catURL
		}
		posts, err := fetchSyllabusItems(client, catURL)
		if err != nil {
			panic(err)
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
