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

// ModelsInspectionChoice models inspection choice
//
// swagger:model models.InspectionChoice
type ModelsInspectionChoice struct {

	// answer
	Answer string `json:"answer,omitempty"`

	// instruction
	Instruction string `json:"instruction,omitempty"`

	// status
	// Required: true
	// Maximum: 4
	// Minimum: 0
	Status *int32 `json:"status"`
}

// Validate validates this models inspection choice
func (m *ModelsInspectionChoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsInspectionChoice) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.MinimumInt("status", "body", int64(*m.Status), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(*m.Status), 4, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models inspection choice based on context it is used
func (m *ModelsInspectionChoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsInspectionChoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsInspectionChoice) UnmarshalBinary(b []byte) error {
	var res ModelsInspectionChoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
