package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/elimity-com/scim"
	"github.com/elimity-com/scim/optional"
	"github.com/elimity-com/scim/schema"
)

type Street struct {
	Name   string
	Number int
}

type Address struct {
	Street   Street
	PostCode string
}

type People map[string][]Address

func main() {
	address := "0.0.0.0:" + os.Getenv("PORT")

	// define resources that our SCIM server can handle. Users is obligatory for Okta - it sends connection test request to /Users
	resourceTypes := []scim.ResourceType{
		{
			ID:          optional.NewString("User"),
			Name:        "User",
			Endpoint:    "/Users",
			Description: optional.NewString("User Account"),
			Schema:      schema.CoreUserSchema(),
			Handler:     NewInMemoryResourceHandler(makeSampleUsers()),
		},
		{
			ID:          optional.NewString("Group"),
			Name:        "Group",
			Endpoint:    "/Groups",
			Description: optional.NewString("User Group"),
			Schema:      schema.CoreGroupSchema(),
			Handler:     NewInMemoryResourceHandler(makeSampleGroups()),
		},
	}

	// define our SCIM server
	server := scim.Server{
		Config: scim.ServiceProviderConfig{
			SupportFiltering: true,
			SupportPatch:     false,
		},
		ResourceTypes: resourceTypes,
	}

	// configure endpoint for SCIM requests; note that requests passed to server.ServeHTTP MUST NOT be prefixed
	http.HandleFunc("/scim/", logRequest(http.StripPrefix("/scim", server).ServeHTTP))
	http.HandleFunc("/auth/scim/", logRequest(http.StripPrefix("/auth/scim", server).ServeHTTP))

	// run
	log.Println("listening at", address)
	http.ListenAndServe(address, nil)
}

// for fields reference, see: schema.CoreUserSchema()
func makeSampleUsers() []scim.Resource {
	attribs1 := map[string]any{
		"userName": "John Doe",
		"active":   true,
		"name": map[string]any{
			"familyName": "Doe",
			"givenName":  "John",
		},
		"emails": []map[string]any{
			{
				"value":   "john_doe@gmail.com",
				"type":    "home",
				"primary": true,
			},
			{
				"value":   "john_doe@acme.com",
				"type":    "work",
				"primary": false,
			},
		},
	}

	attribs2 := map[string]any{
		"userName": "Johny Bravo",
		"active":   true,
		"name": map[string]any{
			"familyName": "Bravo",
			"givenName":  "John",
		},
		"emails": []map[string]any{{
			"value": "johny@gmail.com",
		}},
	}

	attribs3 := map[string]any{
		"userName": "Red Hood",
		"active":   false,
		"name": map[string]any{
			"familyName": "Hood",
			"givenName":  "Red",
		},
		"emails": []map[string]any{{
			"value": "red_hod@gmail.com",
		}},
	}
	return []scim.Resource{
		{ID: "sample-user1", Attributes: attribs1},
		{ID: "sample-user2", Attributes: attribs2},
		{ID: "sample-user3", Attributes: attribs3},
	}
}

// for fields reference, see: schema.CoreGroupSchema()
func makeSampleGroups() []scim.Resource {
	attribs1 := map[string]interface{}{"displayName": "SampleGroup"}
	return []scim.Resource{
		{ID: "sample-group", Attributes: attribs1},
	}
}

// logRequest puts incoming http request into logs and hands over request processing to next function
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	const format = "%-6s %-25s %-25s %s"
	var header = []any{"METHOD", "PATH", "QUERY", "AUTHORIZATION"}
	var once sync.Once

	return func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() { log.Printf(format, header...) })
		log.Printf(format, r.Method, r.URL.Path, r.URL.RawQuery, r.Header.Get("Authorization"))
		next(w, r)
	}
}
