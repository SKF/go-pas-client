// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProblemsRouteNotFoundProblem problems route not found problem
//
// swagger:model problems.RouteNotFoundProblem
type ProblemsRouteNotFoundProblem struct {

	// CorrelationID, an unique identifier for tracing this issue in server logs.
	CorrelationID string `json:"correlationId,omitempty"`

	// Detail, a human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty"`

	// Instance, a URI reference that identifies the specific resource on which the problem occurred.
	Instance string `json:"instance,omitempty"`

	// Status, the HTTP status code associated with this problem occurrence.
	Status int64 `json:"status,omitempty"`

	// Title, a short, human-readable summary of the problem type.
	// This should always be the same value for the same Type.
	Title string `json:"title,omitempty"`

	// Type, a URI reference that identifies the problem type.
	// When dereferenced this should provide human-readable documentation for the
	// problem type. When member is not present it is assumed to be "about:blank".
	Type string `json:"type,omitempty"`
}

// Validate validates this problems route not found problem
func (m *ProblemsRouteNotFoundProblem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this problems route not found problem based on context it is used
func (m *ProblemsRouteNotFoundProblem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProblemsRouteNotFoundProblem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemsRouteNotFoundProblem) UnmarshalBinary(b []byte) error {
	var res ProblemsRouteNotFoundProblem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
