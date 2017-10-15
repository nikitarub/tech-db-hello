// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// FindPetsByStatusURL generates an URL for the find pets by status operation
type FindPetsByStatusURL struct {
	Status []string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindPetsByStatusURL) WithBasePath(bp string) *FindPetsByStatusURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindPetsByStatusURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindPetsByStatusURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/pets/findByStatus"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var statusIR []string
	for _, statusI := range o.Status {
		statusIS := statusI
		if statusIS != "" {
			statusIR = append(statusIR, statusIS)
		}
	}

	status := swag.JoinByFormat(statusIR, "multi")

	for _, qsv := range status {
		qs.Add("status", qsv)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindPetsByStatusURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindPetsByStatusURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindPetsByStatusURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindPetsByStatusURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindPetsByStatusURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindPetsByStatusURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
