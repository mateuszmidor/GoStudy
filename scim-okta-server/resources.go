package main

import (
	"fmt"
	"net/http"

	"github.com/elimity-com/scim"
	scim_errors "github.com/elimity-com/scim/errors"
	"github.com/mateuszmidor/GoStudy/scim/filter"
	orderedmap "github.com/wk8/go-ordered-map/v2"
)

// InMemoryResourceHandler is an in-memory implementation of scim.ResourceHandler interface.
type InMemoryResourceHandler struct {
	resourceName string                                        // e.g. "user", "group"
	resources    *orderedmap.OrderedMap[string, scim.Resource] // keyed by resource ID, orderedmap ensures consistent GetAll results
}

// NewInMemoryResourceHandler creates a new instance of MockResourceHandler.
func NewInMemoryResourceHandler(resourceName string, initResources []scim.Resource) *InMemoryResourceHandler {
	resources := orderedmap.New[string, scim.Resource]()
	for _, r := range initResources {
		resources.Set(r.ID, r)
	}
	return &InMemoryResourceHandler{
		resourceName: resourceName,
		resources:    resources,
	}
}

// Create stores the given attributes and returns a resource with the attributes that are stored and a unique identifier.
func (m *InMemoryResourceHandler) Create(_ *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	if m.isDuplicate(attributes) {
		return scim.Resource{}, scim_errors.ScimErrorUniqueness
	}

	id := fmt.Sprintf("%s%d", m.resourceName, m.resources.Len()+1)
	resource := scim.Resource{
		ID:         id,
		Attributes: attributes,
	}
	m.resources.Set(id, resource)
	return resource, nil
}

// Get returns the resource corresponding to the given identifier.
func (m *InMemoryResourceHandler) Get(_ *http.Request, id string) (scim.Resource, error) {
	if resource, ok := m.resources.Get(id); ok {
		return resource, nil
	}

	return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
}

// GetAll returns a paginated list of resources.
func (m *InMemoryResourceHandler) GetAll(_ *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	// sanitize input
	if params.StartIndex < 1 {
		params.StartIndex = 1
	}

	// handle special case - for Count <= 0 return only the total number of resources and no actual resources
	if params.Count <= 0 {
		return scim.Page{Resources: nil, TotalResults: m.resources.Len()}, nil
	}

	// collect resources matching requested criteria
	var index int // 1-based item index
	var results []scim.Resource
	for pair := m.resources.Oldest(); pair != nil; pair = pair.Next() {
		index++
		if index < params.StartIndex {
			continue
		}
		if len(results) == params.Count {
			break
		}
		match, err := filter.EvalExpression(pair.Value.Attributes, params.Filter)
		if err != nil {
			return scim.Page{}, scim_errors.ScimErrorBadRequest(err.Error())
		}
		if !match {
			continue
		}

		results = append(results, pair.Value)
	}
	return scim.Page{
		Resources:    results,
		TotalResults: len(results),
	}, nil
}

// Replace replaces ALL existing attributes of the resource with the given identifier.
func (m *InMemoryResourceHandler) Replace(_ *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	if _, ok := m.resources.Get(id); !ok {
		return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
	}
	resource := scim.Resource{
		ID:         id,
		Attributes: attributes,
	}
	m.resources.Set(id, resource)
	return resource, nil
}

// Delete removes the resource with the corresponding ID.
func (m *InMemoryResourceHandler) Delete(_ *http.Request, id string) error {
	if _, ok := m.resources.Get(id); !ok {
		return scim_errors.ScimErrorResourceNotFound(id)
	}
	m.resources.Delete(id) // it's ok if resource not present
	return nil
}

// Patch updates one or more attributes of a SCIM resource using a sequence of operations.
func (m *InMemoryResourceHandler) Patch(_ *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	return scim.Resource{}, scim_errors.ScimError{Detail: "Patch is not implemented", Status: http.StatusNotImplemented}
}

func (m *InMemoryResourceHandler) isDuplicate(attributes scim.ResourceAttributes) bool {
	for pair := m.resources.Oldest(); pair != nil; pair = pair.Next() {
		// only userName for User resource needs to be unique, according to schema.CoreUserSchema() and schema.CoreGroupSchema()
		if pair.Value.Attributes["userName"] != nil && pair.Value.Attributes["userName"] == attributes["userName"] {
			return true
		}
	}
	return false
}
