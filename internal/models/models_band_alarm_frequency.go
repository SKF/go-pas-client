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

// ModelsBandAlarmFrequency models band alarm frequency
//
// swagger:model models.BandAlarmFrequency
type ModelsBandAlarmFrequency struct {

	// value
	// Example: 100
	// Required: true
	Value *float64 `json:"value"`

	// The type values are available [here](/v1/docs/service#band-alarm-frequency-value-type).
	// Example: 1
	// Required: true
	// Maximum: 2
	// Minimum: 1
	ValueType *int32 `json:"valueType"`
}

// Validate validates this models band alarm frequency
func (m *ModelsBandAlarmFrequency) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsBandAlarmFrequency) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ModelsBandAlarmFrequency) validateValueType(formats strfmt.Registry) error {

	if err := validate.Required("valueType", "body", m.ValueType); err != nil {
		return err
	}

	if err := validate.MinimumInt("valueType", "body", int64(*m.ValueType), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("valueType", "body", int64(*m.ValueType), 2, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this models band alarm frequency based on context it is used
func (m *ModelsBandAlarmFrequency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelsBandAlarmFrequency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsBandAlarmFrequency) UnmarshalBinary(b []byte) error {
	var res ModelsBandAlarmFrequency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}