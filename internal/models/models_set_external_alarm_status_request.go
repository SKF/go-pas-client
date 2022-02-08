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

// ModelsSetExternalAlarmStatusRequest models set external alarm status request
//
// swagger:model models.SetExternalAlarmStatusRequest
type ModelsSetExternalAlarmStatusRequest struct {

	// This value is optional
	// Example: 123e4567-e89b-12d3-a456-426614174000
	// Format: uuid
	SetBy *strfmt.UUID `json:"setBy,omitempty"`

	// The type values are available [here](/v1/docs/service#alarm-status-type)
	// Example: 2
	// Required: true
	// Maximum: 4
	// Minimum: 0
	Status *int32 `json:"status"`
}

// Validate validates this models set external alarm status request
func (m *ModelsSetExternalAlarmStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSetBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsSetExternalAlarmStatusRequest) validateSetBy(formats strfmt.Registry) error {
	if swag.IsZero(m.SetBy) { // not required
		return nil
	}

	if err := validate.FormatOf("setBy", "body", "uuid", m.SetBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelsSetExternalAlarmStatusRequest) validateStatus(formats strfmt.Registry) error {

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

// ContextValidate validates this models set external alarm status request based on context it is used
func (m *ModelsSetExternalAlarmStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsSetExternalAlarmStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsSetExternalAlarmStatusRequest) UnmarshalBinary(b []byte) error {
	var res ModelsSetExternalAlarmStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
