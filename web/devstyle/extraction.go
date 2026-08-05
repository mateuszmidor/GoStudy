package main

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// SyllabusItem holds the title and relative link of a syllabus entry
// (used for both categories and posts within a category).
type SyllabusItem struct {
	Title string
	Link  string
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
