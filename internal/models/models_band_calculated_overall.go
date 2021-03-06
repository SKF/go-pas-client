// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsBandCalculatedOverall models band calculated overall
//
// swagger:model models.BandCalculatedOverall
type ModelsBandCalculatedOverall struct {

	// unit
	// Example: gE
	Unit string `json:"unit,omitempty"`

	// value
	// Example: 3.14
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this models band calculated overall
func (m *ModelsBandCalculatedOverall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsBandCalculatedOverall) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models band calculated overall based on context it is used
func (m *ModelsBandCalculatedOverall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsBandCalculatedOverall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsBandCalculatedOverall) UnmarshalBinary(b []byte) error {
	var res ModelsBandCalculatedOverall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
