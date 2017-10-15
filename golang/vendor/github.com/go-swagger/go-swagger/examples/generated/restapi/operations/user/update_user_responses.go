// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateUserBadRequestCode is the HTTP code returned for type UpdateUserBadRequest
const UpdateUserBadRequestCode int = 400

/*UpdateUserBadRequest Invalid user supplied

swagger:response updateUserBadRequest
*/
type UpdateUserBadRequest struct {
}

// NewUpdateUserBadRequest creates UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

// WriteResponse to the client
func (o *UpdateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// UpdateUserNotFoundCode is the HTTP code returned for type UpdateUserNotFound
const UpdateUserNotFoundCode int = 404

/*UpdateUserNotFound User not found

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {
}

// NewUpdateUserNotFound creates UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

// WriteResponse to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
