// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostProductsCreatedCode is the HTTP code returned for type PostProductsCreated
const PostProductsCreatedCode int = 201

/*PostProductsCreated Insert product success

swagger:response postProductsCreated
*/
type PostProductsCreated struct {
}

// NewPostProductsCreated creates PostProductsCreated with default headers values
func NewPostProductsCreated() *PostProductsCreated {

	return &PostProductsCreated{}
}

// WriteResponse to the client
func (o *PostProductsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}
