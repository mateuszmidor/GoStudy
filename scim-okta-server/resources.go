package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elimity-com/scim"
	scim_errors "github.com/elimity-com/scim/errors"
)

// InMemoryResourceHandler is an in-memory implementation of scim.ResourceHandler interface.
type InMemoryResourceHandler struct {
	resources map[string]scim.Resource // keyed by resource ID
}

// NewMockResourceHandler creates a new instance of MockResourceHandler.
func NewInMemoryResourceHandler(initResources []scim.Resource) *InMemoryResourceHandler {
	resources := map[string]scim.Resource{}
	for _, r := range initResources {
		resources[r.ID] = r
	}
	return &InMemoryResourceHandler{
		resources: resources,
	}
}

// Create stores the given attributes and returns a resource with the attributes that are stored and a unique identifier.
func (m *InMemoryResourceHandler) Create(r *http.Request, attributes scim.ResourceAttributes) (scim.Resource, error) {
	log.Printf("%+v", attributes)

	if m.isDuplicate(attributes) {
		return scim.Resource{}, scim_errors.ScimErrorUniqueness
	}

	id := fmt.Sprintf("resource%d", len(m.resources)+1)
	resource := scim.Resource{
		ID:         id,
		Attributes: attributes,
	}
	m.resources[id] = resource
	return resource, nil
}

// Get returns the resource corresponding to the given identifier.
func (m *InMemoryResourceHandler) Get(r *http.Request, id string) (scim.Resource, error) {
	if resource, ok := m.resources[id]; ok {
		return resource, nil
	}

	return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
}

// GetAll returns a paginated list of resources.
func (m *InMemoryResourceHandler) GetAll(r *http.Request, params scim.ListRequestParams) (scim.Page, error) {
	// sanity check
	if params.StartIndex < 1 {
		params.StartIndex = 1
	}
	numResoures := len(m.resources)
	if params.StartIndex > numResoures {
		msg := fmt.Sprintf("StartIndex (%d) > total item count (%d)", params.StartIndex, numResoures)
		return scim.Page{}, scim_errors.ScimErrorBadParams([]string{msg})
	}

	// handle special case
	if params.Count <= 0 {
		return scim.Page{Resources: nil, TotalResults: numResoures}, nil
	}

	// filter matching resources
	first := params.StartIndex
	last := first + params.Count // TODO: this is probably wrong as we are filtering the items
	if last > numResoures {
		last = numResoures
	}
	var resources []scim.Resource
	var i int // 1-based index
	for _, resource := range m.resources {
		i++
		if i < first {
			continue
		}
		if i > last {
			break
		}
		// match, err := filter.EvalExpression(resource.Attributes, params.Filter)
		// if err != nil {
		// 	return scim.Page{}, scim_errors.ScimErrorBadRequest(err.Error())
		// }
		// if !match {
		// 	continue
		// }
		resources = append(resources, resource)
	}
	return scim.Page{
		Resources:    resources,
		TotalResults: len(resources),
	}, nil
}

// Replace replaces ALL existing attributes of the resource with the given identifier.
func (m *InMemoryResourceHandler) Replace(r *http.Request, id string, attributes scim.ResourceAttributes) (scim.Resource, error) {
	log.Printf("%s - %+v", id, attributes)
	if _, ok := m.resources[id]; !ok {
		return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
	}
	resource := scim.Resource{
		ID:         id,
		Attributes: attributes,
	}
	m.resources[id] = resource
	return resource, nil
}

// Delete removes the resource with the corresponding ID.
func (m *InMemoryResourceHandler) Delete(r *http.Request, id string) error {
	if _, ok := m.resources[id]; !ok {
		return scim_errors.ScimErrorResourceNotFound(id)
	}
	delete(m.resources, id)
	return nil
}

// Patch updates one or more attributes of a SCIM resource using a sequence of operations.
func (m *InMemoryResourceHandler) Patch(r *http.Request, id string, operations []scim.PatchOperation) (scim.Resource, error) {
	log.Printf("%s - %+v", id, operations)
	return scim.Resource{}, scim_errors.ScimError{Detail: "Patch is not implemented", Status: http.StatusNotImplemented}

	// if _, ok := m.resources[id]; !ok {
	// 	return scim.Resource{}, scim_errors.ScimErrorResourceNotFound(id)
	// }

	// resource := m.resources[id]
	// for _, op := range operations {
	// 	switch op.Op {
	// 	case "add", "replace":
	// 		resource.Attributes[op.Path.AttributePath.String()] = op.Value
	// 	case "remove":
	// 		delete(resource.Attributes, op.Path.String())
	// 	}
	// }

	// m.resources[id] = resource
	// return resource, nil
}

func (m *InMemoryResourceHandler) isDuplicate(attributes scim.ResourceAttributes) bool {
	for _, r := range m.resources {
		// only userName for User resource needs to be unique, according to schema.CoreUserSchema()
		if attributes["userName"] != nil && r.Attributes["userName"] == attributes["userName"] {
			return true
		}
	}
	return false
}
