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

// newAuthorizedClient logs in to a Kajabi/Rails site and returns an
// *http.Client whose cookie jar holds the authenticated session.
// It performs a GET on loginURL to obtain the CSRF token, then POSTs
// the credentials. An error is returned if either request fails.
func newAuthorizedClient(loginURL, email, password string) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Jar: jar}

	// GET the login page to pick up session cookies and the CSRF token.
	resp, err := doGet(client, loginURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("GET login page status:", resp.Status)

	csrfToken := extractCSRFToken(resp.Body)
	fmt.Println("CSRF token found:", csrfToken != "")

	// Build the login form payload.
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

	// POST credentials to log in.
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", loginURL)

	loginResp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer loginResp.Body.Close()
	fmt.Println("Login response status:", loginResp.Status)
	fmt.Println("Login final URL:", loginResp.Request.URL.String())

	return client, nil
}
